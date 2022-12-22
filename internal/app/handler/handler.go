package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"metallplace/internal/app/model"
	"net/http"
	"reflect"
	"time"
)

type IService interface {
	AddMaterialProperty(ctx context.Context, materialSourceId, propertyId int) error
	AddValue(ctx context.Context, materialSourceId int,
		propertyName string, valueFloat float64, valueStr string, createdOn time.Time) error
	AddUniqueMaterial(ctx context.Context, materialName string, sourceName string, materialMarket string, materialUnit string) (int, error)
	InitialImport(ctx context.Context) error
	ParseBook(byte []byte) (model.ChartRaw, error)
	GetMaterialList(ctx context.Context) ([]model.MaterialShortInfo, error)
	GetMaterialValueForPeriod(ctx context.Context, materialSourceId, propertyId int, start string, finish string) ([]model.Price, float64, error)
	GetMaterialSourceInfo(ctx context.Context, id int) (model.MaterialShortInfo, error)
	GetNLastValues(ctx context.Context, materialSourceId, propertyId int, nValues int, finish string) ([]model.Price, error)

	GetChart(ctx context.Context, chartPack model.ChartPack) ([]byte, error)
	GetCachedChart(ctx context.Context, chartPack model.ChartPack) ([]byte, error)
	GetChartRaw(chartRaw model.ChartRaw) ([]byte, error)

	GetReport(repType string, date string) ([]byte, error)

	GetCachedReport(repType string, date string) ([]byte, error)

	GetPropertyList(ctx context.Context, materialSourceId int) ([]model.PropertyShortInfo, error)
	GetPropertyName(ctx context.Context, id int) (string, error)
	AddPropertyToMaterial(ctx context.Context, materialSourceId int, propertyName string, kind string) error
}

type Handler struct {
	service IService
}

func New(srv IService) *Handler {
	return &Handler{service: srv}
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

// Unmarshal request, do work fn(), then marshall response into JSON anf return
func handle[REQ any, RESP any](w http.ResponseWriter, r *http.Request, fn func(req REQ) (RESP, error)) {
	var req REQ
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusBadRequest)
		return
	}

	if isNil(req) != true {
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, fmt.Sprintf("%+v", err), http.StatusBadRequest)
			return
		}
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusBadRequest)
		return
	}

	resp, err := fn(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusInternalServerError)
		return
	}
}
