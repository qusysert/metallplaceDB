package main

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"golang.org/x/sync/errgroup"
	"metallplace/internal/app/api"
	"metallplace/internal/app/handler"
	"metallplace/internal/app/mw"
	"metallplace/internal/app/repository"
	"metallplace/internal/app/service"
	"metallplace/internal/pkg/config"
	"metallplace/pkg/chartclient"
	"metallplace/pkg/docxgenclient"
	db "metallplace/pkg/gopkg-db"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
)

var conn db.IClient

//	@title			Metallplace API
//	@version		1.0
//	@description	Swagger documentation fo metallplace service
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Ivan Demchuk
//@contact.url    http://www.swagger.io/support
//	@contact.email	is.demchuk@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath

//	@securityDefinitions.jwt	Bearer
//	@securityScheme				jwt

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	// Creating logger instance
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	// Loading config
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot load cfg")
	}

	// Creating connection to DB
	conn, err = db.New(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot create connection to db")
	}

	// Creating instances and setting inheritance
	lastRequestTime := time.Time{}
	repo := repository.New()
	chart := chartclient.New(cfg.ChartHost, cfg.ChartPort)
	docxgen := docxgenclient.New(cfg.DocxgenHost, cfg.DocxgenPort)
	modifier, err := api.NewModifier(cfg.ModifierHost, cfg.ModifierPort)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot create modifier grpc client")
	}
	srv := service.New(cfg, repo, chart, docxgen, lastRequestTime, modifier)
	hdl := handler.New(srv)

	// Creating error group for internal and external servers
	eg, egCtx := errgroup.WithContext(context.Background())

	// Running external sever (with auth)
	eg.Go(externalServerFn(egCtx, cfg, hdl, srv))
	log.Printf("External Server started on port %s \n", cfg.HttpPort)

	// Running internal server (no auth)
	eg.Go(internalServerFn(egCtx, cfg, hdl))
	log.Printf("Internal Server started on port %s \n", cfg.InternalHttpPort)

	eg.Go(func() error {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		select {
		case <-sig:
			log.Printf("Receive TERM signal")
			return fmt.Errorf("termination")
		case <-egCtx.Done():
			return nil
		}
	})

	err = eg.Wait()
	if err != nil {
		log.Fatal()
	}

}

func externalServerFn(ctx context.Context, cfg config.Config, hdl *handler.Handler, srv *service.Service) func() error {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.SentryDSN,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}

	// Create an instance of sentryhttp
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	externalRouter := mux.NewRouter()

	// Setting timeout for the externalServer
	externalServer := &http.Server{
		Addr:         ":" + cfg.HttpPort,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 600 * time.Second,
		Handler:      c.Handler(externalRouter),
	}

	// Linking addresses and handlers
	for _, rec := range [...]struct {
		route       string
		handler     http.HandlerFunc
		withoutAuth bool
	}{
		{route: "/swagger.json", handler: func(w http.ResponseWriter, r *http.Request) {
			cwd, _ := os.Getwd()
			p := path.Join(cwd, "docs/swagger.json")
			fmt.Println("Path swagger.json: " + p)
			http.ServeFile(w, r, p)
		}, withoutAuth: true},
		{route: "/swagger/{any:.+}", handler: httpSwagger.Handler(httpSwagger.URL("/swagger.json")), withoutAuth: true},
		{route: "/getValueForPeriod", handler: hdl.GetValueForPeriodHandler},
		{route: "/getMonthlyAvgFeed", handler: hdl.GetMonthlyAvgHandler},
		{route: "/getWeeklyAvgFeed", handler: hdl.GetWeeklyAvgHandler},
		{route: "/getMaterialList", handler: hdl.GetMaterialListHandler},
		{route: "/addValue", handler: hdl.AddValueHandler},
		{route: "/addUniqueMaterial", handler: hdl.AddUniqueMaterialHandler},
		{route: "/initImport", handler: hdl.InitImportHandler},
		{route: "/getNLastValues", handler: hdl.GetNLastValues},
		{route: "/getChart/{specs}", handler: hdl.GetChartHandler},
		{route: "/getReport/{repType}/{date}", handler: hdl.GetReportHandler},
		{route: "/getShortReport", handler: hdl.GetShortReportHandler},
		{route: "/getPropertyList", handler: hdl.GetPropertyListHandler},
		{route: "/getMaterialInfo", handler: hdl.GetMaterialSourceInfoHandler},
		{route: "/getPropertyName", handler: hdl.GetPropertyNameHandler},
		{route: "/addPropertyToMaterial", handler: hdl.AddPropertyToMaterialHandler},
		{route: "/getUnitList", handler: hdl.GetUnitListHandler},
		{route: "/getDeliveryTypeList", handler: hdl.GetDeliveryTypeListHandler},
		{route: "/getChangeSummary", handler: hdl.GetSummaryHandler},
		{route: "/updateMainFile", handler: hdl.UpdateMainFileHandler},

		{route: "/login", handler: hdl.LoginHandler, withoutAuth: true},
	} {
		var h = rec.handler
		if !rec.withoutAuth {
			h = srv.Authenticate(h)
		}
		externalRouter.HandleFunc(rec.route, sentryHandler.HandleFunc(mw.LoggerMiddleware(DbMiddleware(h))))
	}
	return func() error {
		errCh := make(chan error)
		go func() {
			errCh <- externalServer.ListenAndServe()
		}()
		var err error
		select {
		case serverErr := <-errCh:
			err = serverErr
		case <-ctx.Done():
			err = externalServer.Shutdown(ctx)
		}
		log.Printf("External externalServer finished, error: %w\n", err)
		return err
	}
}

func internalServerFn(ctx context.Context, cfg config.Config, hdl *handler.Handler) func() error {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	internalRouter := mux.NewRouter()

	internalServer := &http.Server{
		Addr:         ":" + cfg.InternalHttpPort,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 600 * time.Second,
		Handler:      c.Handler(internalRouter),
	}

	for _, rec := range [...]struct {
		route   string
		handler http.HandlerFunc
	}{
		{route: "/getValueForPeriod", handler: hdl.GetValueForPeriodHandler},
		{route: "/getMonthlyAvgFeed", handler: hdl.GetMonthlyAvgHandler},
		{route: "/getWeeklyAvgFeed", handler: hdl.GetWeeklyAvgHandler},
		{route: "/getMaterialList", handler: hdl.GetMaterialListHandler},
		{route: "/addValue", handler: hdl.AddValueHandler},
		{route: "/addUniqueMaterial", handler: hdl.AddUniqueMaterialHandler},
		{route: "/initImport", handler: hdl.InitImportHandler},
		{route: "/getNLastValues", handler: hdl.GetNLastValues},
		{route: "/getChart/{specs}", handler: hdl.GetChartHandler},
		{route: "/getReport/{repType}/{date}", handler: hdl.GetReportHandler},
		{route: "/getShortReport", handler: hdl.GetShortReportHandler},
		{route: "/getPropertyList", handler: hdl.GetPropertyListHandler},
		{route: "/getMaterialInfo", handler: hdl.GetMaterialSourceInfoHandler},
		{route: "/getPropertyName", handler: hdl.GetPropertyNameHandler},
		{route: "/addPropertyToMaterial", handler: hdl.AddPropertyToMaterialHandler},
		{route: "/updateMainFile", handler: hdl.UpdateMainFileHandler},
		{route: "/getUnitList", handler: hdl.GetUnitListHandler},
		{route: "/getChangeSummary", handler: hdl.GetSummaryHandler},
		{route: "/getDeliveryTypeList", handler: hdl.GetDeliveryTypeListHandler},
	} {
		internalRouter.HandleFunc(rec.route, mw.LoggerMiddleware(DbMiddleware(rec.handler)))
	}

	return func() error {
		errCh := make(chan error)
		go func() {
			errCh <- internalServer.ListenAndServe()
		}()
		var err error
		select {
		case serverErr := <-errCh:
			err = serverErr
		case <-ctx.Done():
			err = internalServer.Shutdown(ctx)
		}
		log.Printf("Internal externalServer finished, error: %w\n", err)
		return err
	}
}

func DbMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(db.AddToContext(ctx, conn))
		next.ServeHTTP(w, r)
	}
}
