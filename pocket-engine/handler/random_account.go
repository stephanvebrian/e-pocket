package handler

import (
	"encoding/json"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
)

func (h *handler) RandomAccount(w http.ResponseWriter, r *http.Request) {
	context := r.Context()

	// Get userID from query parameters instead of body
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		h.writeError(context, model.ErrorResponseOption{
			Writer:     w,
			Request:    r,
			StatusCode: http.StatusBadRequest,
			Response: model.ErrorResponse{
				Code:    model.InvalidParameter,
				Message: "userID parameter is required",
			},
		})
		return
	}

	// Create request object from query param
	req := handlerModel.RandomAccountRequest{
		UserID: userID,
	}

	response, err := h.accountLogic.RandomAccount(context, req)
	if errorResp, ok := err.(model.ErrorResponse); err != nil && ok {
		h.writeError(context, model.ErrorResponseOption{
			Writer:     w,
			Request:    r,
			StatusCode: errorResp.HTTPCode,
			Response:   errorResp,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
