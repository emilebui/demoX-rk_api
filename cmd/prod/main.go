package main

import (
	"context"
	"github.com/emilebui/demoX-rk_api/internal/handlers"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkecho "github.com/rookie-ninja/rk-echo/boot"
)

// @title Swagger DemoX-rk_api API Interface
// @version 1.0
// @description This is a sample demo API server.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic BasicAuth

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	// Create a new boot instance.
	//swaginit.GenerateSwagger()

	boot := rkboot.NewBoot()

	// Register handler
	echoEntry := rkecho.GetEchoEntry("demoX-rk_api")

	// Bootstrap
	boot.Bootstrap(context.TODO())

	echoEntry.Echo.GET("/v1/greeter", handlers.Greeter)

	boot.WaitForShutdownSig(context.TODO())
}
