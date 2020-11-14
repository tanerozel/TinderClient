package TinderClient

import (
	"encoding/json"
	"errors"
)

type client struct {
	requester *TRequest
}

func NewTinderClient(token string) *client {
	return &client{
		NewTRequest(token),
	}
}

func (t *client) SetHeader(key string, value string) {
	t.requester.SetHeader(key, value)
}

func (t *client) RecsCore() ([]RecsCoreUser, error) {
	var recs RecsCoreResponse

	url := "recs/core"
	b, errs := t.requester.Get(url)
	if errs != nil {
		return recs.Data.Results, errs[0]
	}

	err := json.Unmarshal([]byte(b), &recs)
	if err != nil {
		return recs.Data.Results, err
	}

	if recs.Meta.Status != 200 {
		return recs.Data.Results, errors.New("Error getting Recs Core")
	}

	return recs.Data.Results, nil
}

func (t *client) LikeUser(userId string) ([]RecsCoreUser, error) {
	var recs RecsCoreResponse

	url := "like/" + userId
	b, errs := t.requester.Post(url, "")
	if errs != nil {
		return recs.Data.Results, errs[0]
	}

	err := json.Unmarshal([]byte(b), &recs)
	if err != nil {
		return recs.Data.Results, err
	}

	if recs.Meta.Status != 200 {
		return recs.Data.Results, errors.New("Error getting Recs Core")
	}

	return recs.Data.Results, nil
}
