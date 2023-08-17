package handler

import (
	"net/http"
)

type AddUniqueMaterialRequest struct {
	UId          int    `json:"uid"`
	Name         string `json:"name"`
	Source       string `json:"source"`
	Group        string `json:"group"`
	Market       string `json:"market"`
	DeliveryType string `json:"delivery_type"`
	Unit         string `json:"unit"`
}

type AddUniqueMaterialResponse struct {
	Id int `json:"id"`
}

func (h Handler) AddUniqueMaterialHandler(w http.ResponseWriter, r *http.Request) {
	handle(w, r, func(req AddUniqueMaterialRequest) (AddUniqueMaterialResponse, error) {
		err := h.service.AddUniqueMaterial(r.Context(), req.UId, req.Name, req.Group, req.Source, req.Market, req.Unit, req.DeliveryType)
		if err != nil {
			return AddUniqueMaterialResponse{0}, err
		}
		return AddUniqueMaterialResponse{req.UId}, nil
	})
}
