package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Greeter handler
// @Summary Greeter service
// @Id 1
// @version 1.0
// @produce application/json
// @Success 200 {object} GreeterResponse
// @Router /v1/greeter [get]
func Greeter(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Hello World!"),
	})
}

type GreeterResponse struct {
	Message string
}
