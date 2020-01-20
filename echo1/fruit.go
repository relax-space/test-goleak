package main

import (
	"fmt"
	"net/http"
	"nomni/utils/api"
	"time"

	"github.com/labstack/echo"
)

type FruitApiController struct {
}

// localhost:8080/docs
func (d FruitApiController) Init(g echo.Group) {
	g.GET("", d.GetAll)
}

/*
localhost:8080/fruits
localhost:8080/fruits?name=apple
localhost:8080/fruits?skipCount=0&maxResultCount=2
localhost:8080/fruits?skipCount=0&maxResultCount=2&sortby=store_code&order=desc
*/
func (FruitApiController) GetAll(c echo.Context) error {
	f1 := Fruit{
		Code:  "fruit#1",
		Color: "red",
	}
	f2 := Fruit{
		Code:  "fruit#2",
		Color: "green",
	}
	go func() {
		fmt.Println("start")
		time.Sleep(1 * time.Second)
		fmt.Println("end")
	}()
	// time.Sleep(2 * time.Second)
	items := make([]Fruit, 0)
	items = append(items, f1)
	items = append(items, f2)
	return ReturnApiSucc(c, http.StatusOK, api.ArrayResult{
		TotalCount: 2,
		Items:      items,
	})
}
