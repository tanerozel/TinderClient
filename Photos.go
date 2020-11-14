package TinderClient

type UserPhotos struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	FileName string `json:"fileName"`
}

type InstagramPhotos struct {
	Image     string `json:"image"`
	Thumbnail string `json:"thumbnail"`
	Ts        string `json:"ts"`
	Link      string `json:"link"`
}
