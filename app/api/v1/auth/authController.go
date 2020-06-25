package apiv1auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/water25234/golang-postgres/app/controller/api"
)

func GetAuth(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/plain", []byte("Hello, Justin Home!"))
}

func DeleteAuth(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "Hello World DELETE Justin %s", id)
}

func PostAuth(ctx *gin.Context) {
	uid := ctx.Param("uid")

	ctx.JSON(http.StatusOK, api.GetSuccessResponse(gin.H{
		"ThrottleCount": ctx.MustGet("ThrottleCount"),
		"userId":        uid,
	}))
}

func GetThrottle(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", api.GetSuccessResponse(gin.H{
		"ThrottleCount": ctx.MustGet("ThrottleCount"),
	}))
}
