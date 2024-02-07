package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/JaiiR320/SpotifAI/model"
	"github.com/JaiiR320/SpotifAI/scripts"
	"github.com/JaiiR320/SpotifAI/utils"
	"github.com/JaiiR320/SpotifAI/view"
	"github.com/JaiiR320/SpotifAI/view/components"
	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	// Serve static files
	router.Static("/static", "/static")

	// Site Routes
	router.GET("/", HandleShowHome)
	router.GET("/login", HandleSpotifyAuth)

	router.GET("/callback", HandleSpotifyCallback)

	// API Routes
	router.PUT("/tag", HandleAddTag)
	router.DELETE("/tag", HandleDeleteTag)
	router.GET("/filter", HandleFilter)

	scripts.ParsePlaylist("example.json")

	log.Fatal(router.Start(":3000"))
}

func HandleFilter(c echo.Context) error {
	log.Println("Filtering")

	// do some filtering
	selectedItems := scripts.FilterSongs(model.LikedSongs.Items, model.Tags)

	return utils.Render(c, view.TrackList(selectedItems))
}

func HandleDeleteTag(c echo.Context) error {
	values, err := utils.ParseBody(c.Request())
	if err != nil {
		return err
	}
	// Get values from form data
	tag := values.Get("name")

	// Delete tag from model
	err = utils.DeleteFromSlice(&model.Tags, tag)
	if err != nil {
		return c.HTML(http.StatusNoContent, "Tag not found")
	}

	return c.HTML(http.StatusOK, "")
}

func HandleAddTag(c echo.Context) error {
	tag := c.FormValue("tag")

	model.Tags = append(model.Tags, tag)

	return utils.Render(c, components.Tag(tag))
}

func HandleShowHome(c echo.Context) error {
	return utils.Render(c, view.Home())
}

func HandleSpotifyAuth(c echo.Context) error {
	urlStr := `https://accounts.spotify.com/authorize`

	data := url.Values{}
	data.Set("response_type", "code")
	data.Set("client_id", model.Client_id)
	data.Set("scope", "user-read-private user-library-read")
	data.Set("redirect_uri", "http://localhost:3000/callback")

	urlStr = fmt.Sprintf("%s?%s", urlStr, data.Encode())

	return c.Redirect(http.StatusTemporaryRedirect, urlStr)
}

func HandleSpotifyCallback(c echo.Context) error {
	// get access token from spotify through Authentification
	err := getToken(c)
	if err != nil {
		return err
	}

	// get user info from spotify with token
	err = getUser()
	if err != nil {
		return err
	}

	log.Println(model.CurrentUser.DisplayName, "logged in")
	model.Logged = true

	return c.Redirect(http.StatusFound, "/")
}

func getToken(c echo.Context) error {
	idAndSecret := base64.StdEncoding.EncodeToString([]byte(model.Client_id + ":" + model.Client_secret))

	post := scripts.Post("https://accounts.spotify.com/api/token").
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithHeader("Authorization", "Basic "+idAndSecret).
		WithQuery("grant_type", "authorization_code").
		WithQuery("code", c.QueryParam("code")).
		WithQuery("redirect_uri", "http://localhost:3000/callback").
		WithObject(&model.AccessToken)

	return post.Do()
}

func getUser() error {
	req := scripts.Get("https://api.spotify.com/v1/me").
		WithHeader("Authorization", "Bearer "+model.AccessToken.AccessToken).
		WithObject(&model.CurrentUser)
	return req.Do()
}
