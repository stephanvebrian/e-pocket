package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
)

func (h *handler) ListAccount(w http.ResponseWriter, r *http.Request) {
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
	req := handlerModel.ListAccountRequest{
		UserID: userID,
	}

	err := validate.Struct(&req)
	if errorResp, ok := err.(validator.ValidationErrors); ok {
		// Format validation errors into a detailed message
		validationErrors := make(map[string]string)
		for _, fieldError := range errorResp {
			// Construct user-friendly error messages for each field
			validationErrors[fieldError.Field()] = fieldError.Tag()
		}

		// Write the error response
		h.writeError(context, model.ErrorResponseOption{
			Writer:     w,
			Request:    r,
			StatusCode: http.StatusBadRequest,
			Response: model.ErrorResponse{
				Code:    model.ValidationError,
				Message: "Validation failed",
				Details: validationErrors, // Include field-specific details
			},
		})
		return
	}

	response, err := h.accountLogic.ListAccount(context, req)
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
