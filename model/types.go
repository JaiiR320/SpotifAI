package model

type User struct {
	DisplayName  string `json:"display_name"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href   string `json:"href"`
	Id     string `json:"id"`
	Images []struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

// Playlist types
type Track struct {
	Name  string `json:"name"`
	Album struct {
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
		Images []struct {
			Height int    `json:"height"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name string `json:"name"`
	} `json:"album"`
}

type Item struct {
	Track `json:"track"`
}

type Playlist struct {
	Items []Item `json:"items"`
}

// {
//     "songs":[{
//         "title":"HACIENDO QUE ME AMAS"
//     },{
// 	"title":"AL CALLAO'"
//     }]
// }
