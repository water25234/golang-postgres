package usermodel

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/water25234/golang-postgres/server"
)

type Item struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func GetUserList() (*Item, error) {

	// defer server.DB.Close()

	var item Item
	err := server.DB.QueryRow("SELECT * FROM inventory;").Scan(&item.Id, &item.Name, &item.Quantity)

	if err != nil {
		panic(err)
	}

	fmt.Println(&item)

	return &item, err
}
