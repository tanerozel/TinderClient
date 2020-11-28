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

func LoginSms(phone string) {
	r := NewRequest()
	var headerJson = []byte(`{
    "accept": "application/json",
    "accept-language": "tr,en;q=0.9,en-GB;q=0.8,en-US;q=0.7",
    "app-session-id": "2b5351ba-2bec-49b7-be55-aeaff9b76ea1",
    "app-session-time-elapsed": "-3508674",
    "app-version": "1026300",
    "content-type": "application/x-google-protobuf",
    "funnel-session-id": "73ad6deae5a0b3f6",
    "persistent-device-id": "7e6e7592-d428-4109-9fdc-f022513b7897",
    "platform": "web",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "cross-site",
    "tinder-version": "2.63.0",
    "user-session-id": "null",
    "user-session-time-elapsed": "24159",
    "x-supported-image-formats": "jpeg"
  }`)
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	json.Unmarshal(headerJson, &c)

	r.Post("https://api.gotinder.com/v3/auth/login?locale=tr").Send("\n\u000e\n\u000c" + phone)
	for k, v := range c {
		r.Set(k, string(v))
	}
	r.End()
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

func (t *client) LikeUser(userId string) (string, error) {
	var recs RecsCoreResponse

	url := "like/" + userId
	b, errs := t.requester.Post(url, "")
	if errs != nil {
		return b, errs[0]
	}

	err := json.Unmarshal([]byte(b), &recs)
	if err != nil {
		return b, err
	}

	return b, nil
}

func (t *client) loginAuth(phone string) {
	t.requester.Post("auth/login", phone)
}
