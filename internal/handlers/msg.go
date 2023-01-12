package handlers

import (
	"github.com/emilebui/demoX-rk_api/internal/services"
	"github.com/labstack/echo/v4"
	rkecho "github.com/rookie-ninja/rk-echo/boot"
	"net/http"
)

type MessageHandler struct {
	s *services.Service
}

func NewMessageHandler(s *services.Service) Handler {
	return &MessageHandler{
		s: s,
	}
}

func (h *MessageHandler) RegisterAPI(echoEntry *rkecho.EchoEntry) {
	echoEntry.Echo.POST("/v1/push", h.PushMessage)
	echoEntry.Echo.POST("/v1/push_many", h.PushMessages)
}

// PushMessage handler
// @Summary Push a message to a queue
// @Description Send a message to a queue (redis or kafka)
// @version 1.0
// @Tag Message
// @produce application/json
// @Param input body ResAndInput true "Input required"
// @Success 200 {object} ResAndInput
// @Router /v1/push [post]
func (h *MessageHandler) PushMessage(ctx echo.Context) error {

	input := new(ResAndInput)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.s.Process(input.Message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, ResAndInput{Message: "Push message to queue successfully!!"})
}

type ResAndInput struct {
	Message string `json:"message"`
}

type PushMessagesInput struct {
	Messages []string `json:"messages"`
}

// PushMessages handler
// @Summary Push many messages to a queue
// @Description Send multiple messages to a queue (redis or kafka)
// @version 1.0
// @Tag Message
// @produce application/json
// @Param input body PushMessagesInput true "Input required"
// @Success 200 {object} ResAndInput
// @Router /v1/push_many [post]
func (h *MessageHandler) PushMessages(ctx echo.Context) error {

	input := new(PushMessagesInput)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, message := range input.Messages {
		err := h.s.Process(message)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	return ctx.JSON(http.StatusOK, ResAndInput{Message: "Push messages to queue successfully!!"})
}
