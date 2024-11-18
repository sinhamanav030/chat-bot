package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	messages map[string]string
	counter  int
}

func NewHandler() Handler {
	mssgs := make(map[string]string)
	return Handler{
		messages: mssgs,
		counter:  0,
	}
}

type PostMessageRequest struct {
	Message string `json:"message"`
}

type PostMessageResponse struct {
	Status    string `json:"status,omitempty"`
	Error     string `json:"error,omitempty"`
	MessageId string `json:"message_id,omitempty"`
}

type GetMessageResponse struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func (h *Handler) PostMessageHandler(ctx *gin.Context) {
	var req PostMessageRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, PostMessageResponse{
			Error: err.Error(),
		})
		return
	}

	key := fmt.Sprintf("%d", h.counter)

	h.messages[key] = req.Message
	h.counter = h.counter + 1

	ctx.JSON(http.StatusCreated, PostMessageResponse{
		Status:    "ok",
		MessageId: key,
	})
}

func (h *Handler) GetMessageHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	mssg, ok := h.messages[id]

	if !ok {
		ctx.JSON(http.StatusBadRequest, GetMessageResponse{
			Error: "cannot find mssg for given mssg id",
		})
		return
	}

	ctx.JSON(http.StatusOK, GetMessageResponse{
		Message: mssg,
	})
}
