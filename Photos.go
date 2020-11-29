package TinderClient

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

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

func DownloadImage(photo UserPhotos, path string) {
	context.Background()
	//redis.NewClient(redis.Options{DB: })
	filePath := path + "/" + photo.FileName
	_, fileErr := os.Stat(filePath)
	if !os.IsNotExist(fileErr) {
		fmt.Println(filePath + " is dir")
	} else {
		response, e := http.Get(photo.URL)
		if e != nil {
			log.Fatal(e)
		}

		defer response.Body.Close()

		file, err := os.Create(filePath)

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(filePath + " dowland")
	}

}
