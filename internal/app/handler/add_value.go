package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type AddValueRequest struct {
	MaterialSourceId string    `json:"material_source_id"`
	PropertyName     string    `json:"property_name"`
	ValueFloat       string    `json:"value_float"`
	ValueStr         string    `json:"value_str"`
	CreatedOn        time.Time `json:"created_on"`
}

type AddValueResponse struct {
	Success bool
}

func (h Handler) AddValueHandler(w http.ResponseWriter, r *http.Request) {
	handle(w, r, func(req AddValueRequest) (AddValueResponse, error) {
		var valueFloat float64
		var err error

		if req.ValueFloat != "" {
			valueFloat, err = strconv.ParseFloat(req.ValueFloat, 64)
		}

		MaterialSourceId, err := strconv.Atoi(req.MaterialSourceId)
		if err != nil {
			return AddValueResponse{false}, fmt.Errorf("cant convert material_source_id to int: %v", err)
		}

		if err != nil {
			return AddValueResponse{false}, err
		}
		err = h.service.AddValue(r.Context(), MaterialSourceId,
			req.PropertyName, valueFloat, req.ValueStr, req.CreatedOn)
		if err != nil {
			return AddValueResponse{false}, err
		}
		return AddValueResponse{true}, err
	})
}
