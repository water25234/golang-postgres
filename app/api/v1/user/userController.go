package apiv1user

import (
	"net/http"

	"github.com/water25234/golang-postgres/app/controller/api"
	"github.com/water25234/golang-postgres/core/log"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	log.Info("test get user log")
	ctx.JSON(http.StatusOK, api.GetSuccessResponse(gin.H{"userId": uid}))
}
