package main

import (
	"github.com/emilebui/demoX-rk_api/cmd"
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
	cmd.Run()
}
