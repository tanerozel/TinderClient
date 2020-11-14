package TinderClient

import "time"

type Instagram struct {
	LastFetchTime         time.Time         `json:"last_fetch_time"`
	CompletedInitialFetch bool              `json:"completed_initial_fetch"`
	Photos                []InstagramPhotos `json:"photos"`
	MediaCount            int               `json:"media_count"`
	ProfilePicture        string            `json:"profile_picture"`
	Username              string            `json:"username"`
}
