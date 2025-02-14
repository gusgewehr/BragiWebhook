package controller

import (
	"BragiWebhooks/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SubscribeController struct {
	Env *infrastructure.Env
}

func (sc *SubscribeController) Subscribe(ctx *gin.Context) {
	hubMode := ctx.Query("hub.mode")
	hubToken := ctx.Query("hub.verify_token")

	if hubMode == "subscribe" || hubToken == sc.Env.SubscribeToken {
		hubChallenge := ctx.Query("hub.challenge")
		challenge, _ := strconv.Atoi(hubChallenge)
		ctx.JSON(http.StatusOK, challenge)
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subscribe attempt"})
	return
}
