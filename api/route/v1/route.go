package route

import (
	"log"
	"time"

	"github.com/astertechs-dev/bizportal-go-backend/bootstrap"
	"github.com/astertechs-dev/bizportal-go-backend/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, routerV1 *gin.RouterGroup) {
	log.Println(routerV1)
}
