package TinderClient

import "time"

type User struct {
	Bio           string        `json:"bio"`
	BirthDate     time.Time     `json:"birth_date"`
	Name          string        `json:"name"`
	PingTime      time.Time     `json:"ping_time"`
	Photos        []UserPhotos  `json:"photos"`
	Jobs          []interface{} `json:"jobs"`
	Schools       []interface{} `json:"schools"`
	ID            string        `json:"_id"`
	Gender        int           `json:"gender"`
	BirthDateInfo string        `json:"birth_date_info"`
}
