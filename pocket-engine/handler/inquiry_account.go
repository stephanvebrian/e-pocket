package handler

import (
	"encoding/json"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
)

func (h *handler) InquiryAccount(w http.ResponseWriter, r *http.Request) {
	context := r.Context()

	accountNumber := r.URL.Query().Get("accountNumber")
	if accountNumber == "" {
		h.writeError(context, model.ErrorResponseOption{
			Writer:     w,
			Request:    r,
			StatusCode: http.StatusBadRequest,
			Response: model.ErrorResponse{
				Code:    model.InvalidParameter,
				Message: "accountNumber parameter is required",
			},
		})
		return
	}

	req := handlerModel.InquiryAccountRequest{
		AccountNumber: accountNumber,
	}

	response, err := h.accountLogic.Inquiry(context, req)
	if errorResp, ok := err.(model.ErrorResponse); err != nil && ok {
		h.writeError(context, model.ErrorResponseOption{
			Writer:     w,
			Request:    r,
			StatusCode: errorResp.HTTPCode,
			Response:   errorResp,
		})
		return
	}
	// unexpected error
	if err != nil {
		h.writeError(context, model.ErrorResponseOption{
			Writer:     w,
			Request:    r,
			StatusCode: http.StatusInternalServerError,
			Response: model.ErrorResponse{
				Code:    model.UnexpectedError,
				Message: "Failed to create transfer",
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
