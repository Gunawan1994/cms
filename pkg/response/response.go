package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data"`
}

func SetResponse(ctx echo.Context, httpstatus int, msg string, meta interface{}, data interface{}) error {
	return ctx.JSON(httpstatus, Response{
		Message: msg,
		Meta:    meta,
		Data:    data,
	})
}
