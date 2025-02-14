package controller

import (
	"BragiWebhooks/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebhookController struct {
	ReceivedTextMessageUC domain.ReceivedMessageUseCase
}

func (wc *WebhookController) Webhook(ctx *gin.Context) {
	var receivedTextMessage domain.ReceivedMessage

	err := ctx.ShouldBind(&receivedTextMessage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = wc.ReceivedTextMessageUC.Send(ctx, receivedTextMessage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": receivedTextMessage})
}
