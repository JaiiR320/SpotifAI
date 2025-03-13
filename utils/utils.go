package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// added a test comment
func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func WriteJSON(c echo.Context, status int, v any) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(status)
	return json.NewEncoder(c.Response()).Encode(v)
}

func DeleteFromSlice(slice *[]string, item string) error {
	for i, v := range *slice {
		if v == item {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

func ParseBody(req *http.Request) (url.Values, error) {
	// Read the body
	body := req.Body
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, body); err != nil {
		return nil, err
	}
	bodyStr := buf.String()

	// Parse the body as form data
	values, err := url.ParseQuery(bodyStr)
	if err != nil {
		return nil, err
	}

	return values, nil
}
