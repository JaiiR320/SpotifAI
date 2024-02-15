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
	"github.com/JaiiR320/SpotifAI/view/layout"
	"github.com/JaiiR320/SpotifAI/view/pages"
	"github.com/imroc/req/v3"

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
	router.DELETE("/tag/:name", HandleDeleteTag)

	log.Fatal(router.Start(":3000"))
}

func HandleDeleteTag(c echo.Context) error {
	tag := c.Param("name")
	err := model.DeleteTag(tag)
	if err != nil {
		return c.HTML(http.StatusNoContent, "Tag not found")
	}

	log.Println("Deleted tag:", tag)

	model.FilteredSongs = scripts.FilterSongs(model.LikedSongs, model.Tags)

	return utils.Render(c, layout.Content())
}

func HandleAddTag(c echo.Context) error {
	tag := c.FormValue("tag")

	model.AddTag(tag)
	log.Println("Added tag:", tag)

	model.FilteredSongs = scripts.FilterSongs(model.LikedSongs, model.Tags)

	return utils.Render(c, layout.Content())
}

func HandleShowHome(c echo.Context) error {
	return utils.Render(c, pages.Home())
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
	err = getLikedSongs()
	if err != nil {
		return err
	}

	model.Logged = true

	return c.Redirect(http.StatusFound, "/")
}

func getLikedSongs() error {
	log.Println("Getting liked songs")
	client := req.C()
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+model.AccessToken.AccessToken).
		SetSuccessResult(&model.LikedSongsResponse).
		Get("https://api.spotify.com/v1/me/tracks")
	if err != nil {
		return err
	}

	if resp.GetStatusCode() != http.StatusOK {
		return fmt.Errorf("no response from Spotify")
	}

	model.LikedSongs = model.LikedSongsResponse.Items
	model.FilteredSongs = model.LikedSongsResponse.Items

	return nil
}

func getToken(c echo.Context) error {
	log.Println("Getting token")
	idAndSecret := base64.StdEncoding.EncodeToString([]byte(model.Client_id + ":" + model.Client_secret))

	client := req.C()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Authorization", "Basic "+idAndSecret).
		SetQueryParam("grant_type", "authorization_code").
		SetQueryParam("code", c.QueryParam("code")).
		SetQueryParam("redirect_uri", "http://localhost:3000/callback").
		SetSuccessResult(&model.AccessToken).
		Post("https://accounts.spotify.com/api/token")
	if err != nil {
		return err
	}

	if resp.GetStatusCode() != http.StatusOK {
		return fmt.Errorf("no response from Spotify")
	}

	return nil
}

func getUser() error {
	log.Println("Getting user info")
	client := req.C()
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+model.AccessToken.AccessToken).
		SetSuccessResult(&model.CurrentUser).
		Get("https://api.spotify.com/v1/me")
	if err != nil {
		return err
	}

	if resp.GetStatusCode() != http.StatusOK {
		return fmt.Errorf("no response from Spotify")
	}

	return nil
}
