package main

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func WriteJSON(c echo.Context, status int, v any) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(status)
	return json.NewEncoder(c.Response()).Encode(v)
}
