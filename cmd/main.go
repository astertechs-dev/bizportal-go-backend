package cmd

import (
	"time"

	routeV1 "github.com/astertechs-dev/bizportal-go-backend/api/route/v1"
	"github.com/astertechs-dev/bizportal-go-backend/bootstrap"
	gin2 "github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin2.Default()

	routerV1 := gin.Group("v1")

	routeV1.Setup(env, timeout, db, routerV1)

	err := gin.Run(env.ServerAddress)
	if err != nil {
		return
	}
}
