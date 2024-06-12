package handler

import (
	"context"
	"encoding/json"
	"image-watermark/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CarUsecase represent the car's Usecases
type Usecase interface {
	AddWatermark(ctx context.Context) (*interface{}, error)
}

type ResponseError struct {
	Message string `json:"message"`
}

type Handler struct {
	Usecase Usecase
}

func (h *Handler) AddWatermark(e echo.Context) error {
	resp, err := h.Usecase.AddWatermark(e.Request().Context())
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			ResponseError{Message: domain.ErrInternalServerError.Error()},
		)
	}

	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	e.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(e.Response()).Encode(resp)

}
