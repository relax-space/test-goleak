package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/test"
)

func TestFruitCRUD(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		test.Assert(t, handleWithFilter != nil, "handleWithFilter is nil")
		req := httptest.NewRequest(echo.GET, "/v1/fruits", nil)
		test.Assert(t, req != nil, "req is nil")
		rec := httptest.NewRecorder()
		test.Assert(t, rec != nil, "rec is nil")
		test.Ok(t, handleWithFilter(FruitApiController{}.GetAll, echoApp.NewContext(req, rec)))
		test.Equals(t, http.StatusOK, rec.Code)

		var v struct {
			Result struct {
				TotalCount int     `json:"totalCount"`
				Items      []Fruit `json:"items"`
			} `json:"result"`
			Success bool `json:"success"`
		}
		test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
		test.Equals(t, v.Result.TotalCount, 2)
		test.Equals(t, v.Result.Items[0].Code, "fruit#1")
		test.Equals(t, v.Result.Items[1].Code, "fruit#2")
		time.Sleep(3*time.Second)
	})
}
