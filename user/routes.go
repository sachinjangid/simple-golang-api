package user

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-golang-api/user/controller"
)

// APIRoutes :
func APIRoutes(r *gin.Engine) {

	r.GET("/", testFunc)

	r.POST("/login", controller.UserLogin)
	r.GET("/user/info", controller.GetUsers)
}

func testFunc(ctx *gin.Context) {
	ctx.Writer.Write([]byte("<div><h1>Simple API</h1></div>"))
}
