package route

import (
	"BragiWebhooks/adapter/controller"
	"BragiWebhooks/infrastructure"
	"BragiWebhooks/repository"
	"BragiWebhooks/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(router *gin.Engine, app *infrastructure.Application) {

	router.Use(cors.Default())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})
	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Method not found"})
	})

	rp := repository.NewExampleRepository(app.Amqp, app.Logger)

	uc := usecase.NewExampleUseCase(app.Logger, rp)

	subCtrl := &controller.SubscribeController{Env: app.Env}

	webhCtrl := &controller.WebhookController{ReceivedTextMessageUC: uc}

	router.GET("/webhook", subCtrl.Subscribe)

	router.POST("/webhook", webhCtrl.Webhook)

}
