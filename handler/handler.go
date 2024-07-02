package handler

import (
	"context"
	"image-watermark/domain"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CarUsecase represent the car's Usecases
type Usecase interface {
	AddWatermark(context.Context, multipart.File, domain.Request) error
}

type ResponseError struct {
	Message string `json:"message"`
}

type Handler struct {
	Usecase Usecase
}

func (h *Handler) AddWatermark(e echo.Context) error {
	var req domain.Request
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(
			http.StatusBadRequest,
			ResponseError{Message: err.Error()},
		)
	}
	req.Authorization = e.Request().Header.Get("Authorization")
	file, err := e.FormFile("file")
	if err != nil {
		return e.JSON(
			http.StatusBadRequest,
			ResponseError{Message: err.Error()},
		)
	}

	src, err := file.Open()
	if err != nil {
		return e.JSON(
			http.StatusBadRequest,
			ResponseError{Message: err.Error()},
		)
	}
	defer src.Close()

	err = h.Usecase.AddWatermark(e.Request().Context(), src, req)
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			ResponseError{Message: err.Error()},
		)
	}

	return e.JSON(201, map[string]string{
		"message": "success",
	})

}
