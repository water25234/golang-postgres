package userserver

import (
	usermodel "github.com/water25234/golang-postgres/app/model"
)

func GetUserList() interface{} {
	item, err := usermodel.GetUserList()

	if err != nil {
		panic(err)
	}

	return item
}
