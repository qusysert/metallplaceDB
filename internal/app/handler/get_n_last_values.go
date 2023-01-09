package handler

import (
	"metallplace/internal/app/model"
	"net/http"
)

type LastValuesRequest struct {
	MaterialSourceId int    `json:"material_source_id" `
	PropertyId       int    `json:"property_id"`
	NValues          int    `json:"n_values"`
	Finish           string `json:"finish"`
}

type LastValuesResponse struct {
	PriceFeed []model.Price `json:"price_feed"`
}

func (h Handler) GetNLastValues(w http.ResponseWriter, r *http.Request) {
	handle(w, r, func(req LastValuesRequest) (LastValuesResponse, error) {
		priceFeed, err := h.service.GetNLastValues(r.Context(), req.MaterialSourceId, req.PropertyId, req.NValues, req.Finish)
		return LastValuesResponse{priceFeed}, err
	})
}
