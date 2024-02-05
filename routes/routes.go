package routes

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()

	r.GET("api/v1/ping", controllers.Ping)

}
