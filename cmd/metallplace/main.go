package main

import (
	"github.com/gorilla/mux"
	"log"
	"metallplace/internal/app/handler"
	"metallplace/internal/app/repository"
	"metallplace/internal/app/service"
	"metallplace/internal/pkg/config"
	"metallplace/internal/pkg/converter"
	"metallplace/pkg/chartclient"
	"metallplace/pkg/gopkg-db"
	"net/http"
	"time"
)

var conn db.IClient

func main() {
	// Loading config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load cfg:", err)
	}

	// Creating connection to DB
	conn, err = db.New(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Creating instances and setting inheritance
	repo := repository.New()
	chart := chartclient.New(cfg.ChartHost, cfg.ChartPort)
	conv := converter.New(cfg.ConvHost, cfg.ConvPort)
	srv := service.New(cfg, repo, chart, conv)
	hdl := handler.New(srv)

	// Setting timeout for the server
	server := &http.Server{
		Addr:         ":" + cfg.HttpPort,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	router := mux.NewRouter()

	// Linking addresses and handlers
	for _, rec := range [...]struct {
		route   string
		handler http.HandlerFunc
	}{
		{route: "/getValueForPeriod", handler: hdl.GetValueForPeriodHandler},
		{route: "/getMaterials", handler: hdl.GetMaterialHandler},
		{route: "/addValue", handler: hdl.AddValueHandler},
		{route: "/addUniqueMaterial", handler: hdl.AddUniqueMaterialHandler},
		{route: "/initImport", handler: hdl.InitImport},
		{route: "/getNLastValues", handler: hdl.GetNLastValues},
		{route: "/getChart/{specs}", handler: hdl.GetChartHandler},
		{route: "/getReport/{date}", handler: hdl.GetReportHandler},
	} {
		router.HandleFunc(rec.route, DbMiddleware(rec.handler))
	}

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web"))))

	http.Handle("/", router)

	log.Printf("Server started on port %s \n", cfg.HttpPort)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func DbMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(db.AddToContext(ctx, conn))
		next.ServeHTTP(w, r)
	}
}
