package router

import (
	"github.com/gin-gonic/gin"
	coreserver "github.com/water25234/golang-postgres/core/server"
	"github.com/water25234/golang-postgres/server"
)

var Router *gin.Engine

func init() {
	coreserver.SetServerGonfig()
	coreserver.SetAppConfig()
	server.InitRedis()
	server.InitDataBase()
}

func StartServer() {
	Router = SetupRouter()
	Router.Run()
}
