package main

import (
	"fmt"
	"os"
	"testing"

	"go.uber.org/goleak"

	"nomni/utils/validator"

	"github.com/labstack/echo"
)

var (
	echoApp          *echo.Echo
	handleWithFilter func(handlerFunc echo.HandlerFunc, c echo.Context) error
)

func TestMain(m *testing.M) {
	fmt.Println(1212)
	enterTest()
	fmt.Println(1414)
	goleak.VerifyTestMain(m)
	code := m.Run()
	os.Exit(code)
}

func enterTest() {

	echoApp = echo.New()
	echoApp.Validator = validator.New()

	header := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c.SetRequest(req)
			return next(c)
		}
	}
	handleWithFilter = func(handlerFunc echo.HandlerFunc, c echo.Context) error {
		return header(handlerFunc)(c)
	}
}
