package main

import (
	"nomni/utils/api"

	"github.com/labstack/echo"
)

func ReturnApiFail(c echo.Context, status int, err error) error {

	return c.JSON(status, api.Result{
		Success: false,
		Error:   api.UnknownError(err),
	})
}

func ReturnApiSucc(c echo.Context, status int, result interface{}) error {

	return c.JSON(status, api.Result{
		Success: true,
		Result:  result,
	})
}
