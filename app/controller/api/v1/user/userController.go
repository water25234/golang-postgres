package apiv1user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/water25234/golang-postgres/app/controller/api"
	userserver "github.com/water25234/golang-postgres/app/server"
	"github.com/water25234/golang-postgres/core/log"
)

func GetUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	log.Info("test get user log")
	ctx.JSON(http.StatusOK, api.GetSuccessResponse(gin.H{"userId": uid}))
}

type User struct {
	FirstName      string `validate:"required"`
	LastName       string `validate:"required"`
	Age            uint8  `validate:"gte=0,lte=130"`
	Email          string `validate:"required,email"`
	FavouriteColor string `validate:"hexcolor|rgb|rgba"`
	Addresses      string `validate:"required"`
}

var validate *validator.Validate

func AddUser(ctx *gin.Context) {
	validate = validator.New()

	user := &User{
		FirstName:      "",
		LastName:       "11",
		Age:            45,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000",
		Addresses:      "justin.huang@kkday.com",
	}

	err := validate.Struct(user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("err is ", err)
			return
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("err is ", err)
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return
	}

	ctx.JSON(http.StatusOK, api.GetSuccessResponse(gin.H{"userId": "123"}))
}

func GetUserList(ctx *gin.Context) {
	uid := ctx.Param("uid")
	log.Info("test get user log")
	ctx.JSON(http.StatusOK, api.GetSuccessResponse(gin.H{"userId": uid}))
}

func GetUserListVersion2(ctx *gin.Context) {
	response := userserver.GetUserList()

	// fmt.Println(response)
	json, err := json.Marshal(response)

	if err != nil {
		ctx.JSON(http.StatusFailedDependency, api.GetSuccessResponse(response))
		return
	}
	value := string(json)

	fmt.Println(value)

	ctx.JSON(http.StatusOK, api.GetSuccessResponse(response))
}
