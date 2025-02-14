package controller

import (
	"BragiWebhooks/infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type SubscribeController struct {
	Env *infrastructure.Env
}

func (sc *SubscribeController) Subscribe(ctx *gin.Context) {
	hubMode := ctx.Query("hub.mode")
	hubToken := ctx.Query("hub.verify_token")

	if hubMode == "subscribe" || hubToken == sc.Env.SubscribeToken {
		hubChallenge := ctx.Query("hub.challenge")
		ctx.JSON(http.StatusOK, fmt.Sprintf("%s", strings.Replace(hubChallenge, "\\\"", "", 2)))
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subscribe attempt"})
	return
}
