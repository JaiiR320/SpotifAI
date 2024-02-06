package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/JaiiR320/SpotifAI/model"
	"github.com/JaiiR320/SpotifAI/view"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var client_id = `04c4fff349274db8b607e7a927f2ca5a`
var client_secret = `a97b3bed058d4a1ea4502654bcf554f6`

func main() {
	router := echo.New()

	router.Static("/static", "/static")

	router.GET("/", HandleShowHome)
	router.GET("/login", HandleSpotifyAuth)

	router.GET("/callback", HandleSpotifyCallback)

	// scripts.ParsePlaylist("example.json")

	log.Fatal(router.Start(":3000"))
}

func HandleShowHome(c echo.Context) error {
	return Render(c, view.Home())
}

func HandleSpotifyAuth(c echo.Context) error {
	urlStr := `https://accounts.spotify.com/authorize`

	data := url.Values{}
	data.Set("response_type", "code")
	data.Set("client_id", client_id)
	data.Set("scope", "user-read-private user-library-read")
	data.Set("redirect_uri", "http://localhost:3000/callback")

	urlStr = fmt.Sprintf("%s?%s", urlStr, data.Encode())

	return c.Redirect(http.StatusTemporaryRedirect, urlStr)
}

func HandleSpotifyCallback(c echo.Context) error {
	err := getToken(c)
	if err != nil {
		return err
	}

	err = getUser()
	if err != nil {
		log.Println("GET USER ", err)
		return err
	}

	c.QueryParams().Del("code")

	return c.Redirect(http.StatusFound, "/")
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func getToken(c echo.Context) error {
	code := c.QueryParam("code")

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", "http://localhost:3000/callback")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+client_id)
	req.Header.Set("Authorization",
		"Basic "+base64.StdEncoding.EncodeToString([]byte(client_id+":"+client_secret)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return err
	}

	model.AccessToken = tokenResponse.AccessToken
	log.Println(model.AccessToken)
	return nil
}

func getUser() error {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+model.AccessToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var user model.User

	err = json.Unmarshal(body, &user)

	if err != nil {
		return err
	}

	model.CurrentUser = user
	return nil
}

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
