package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	rkecho "github.com/rookie-ninja/rk-echo/boot"
	"net/http"
)

type DemoHandler struct {
}

func NewDemoHandler() Handler {
	return &DemoHandler{}
}

func (h *DemoHandler) RegisterAPI(echoEntry *rkecho.EchoEntry) {
	echoEntry.Echo.GET("/v1/greeter", h.Greeter)
}

// Greeter handler
// @Summary Greeter service
// @version 1.0
// @Tag Demo
// @produce application/json
// @Success 200 {object} GreeterResponse
// @Router /v1/greeter [get]
func (h *DemoHandler) Greeter(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Hello World!"),
	})
}

type GreeterResponse struct {
	Message string
}
