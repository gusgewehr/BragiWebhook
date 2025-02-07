package main

import (
	"BragiWebhooks/adapter/route"
	"BragiWebhooks/infrastructure"
	"github.com/gin-gonic/gin"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	envPath := ".env"

	if len(os.Args) > 1 {
		envPath = os.Args[1]
	}

	app := infrastructure.App(envPath)

	httpServer := gin.Default()

	route.SetupRoutes(httpServer, app)

	err := httpServer.Run()

	if err != nil {
		panic(err)
	}

}
