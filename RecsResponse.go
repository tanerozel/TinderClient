package TinderClient

type RecsCoreUser struct {
	DistanceMi      int           `json:"distance_mi"`
	ConnectionCount int           `json:"connection_count"`
	CommonLikes     []interface{} `json:"common_likes"`
	CommonFriends   []interface{} `json:"common_friends"`
	ContentHash     string        `json:"content_hash"`
	Badges          []interface{} `json:"badges,omitempty"`
	Teaser          struct {
		String string `json:"string"`
	} `json:"teaser"`
	Teasers           []interface{} `json:"teasers"`
	SNumber           int           `json:"s_number"`
	SpotifyThemeTrack interface{}   `json:"spotify_theme_track,omitempty"`
	Instagram         Instagram     `json:"instagram,omitempty"`
	CommonConnections []interface{} `json:"common_connections,omitempty"`
	CommonInterests   []interface{} `json:"common_interests,omitempty"`
	UncommonInterests []interface{} `json:"uncommon_interests,omitempty"`
	SpotifyTopArtists []struct {
		Selected bool   `json:"selected"`
		Name     string `json:"name"`
		ID       string `json:"id"`
		TopTrack struct {
			URI   string `json:"uri"`
			Album struct {
				Images []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"images"`
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"album"`
			Name       string `json:"name"`
			PreviewURL string `json:"preview_url"`
			Artists    []struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"artists"`
			ID string `json:"id"`
		} `json:"top_track"`
	} `json:"spotify_top_artists,omitempty"`
	IsTraveling         bool   `json:"is_traveling,omitempty"`
	ShowGenderOnProfile bool   `json:"show_gender_on_profile"`
	Type                string `json:"type"`
	User                User   `json:"user"`
}

type RecsCoreResponse struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
		Results []RecsCoreUser `json:"results"`
	} `json:"data"`
}
