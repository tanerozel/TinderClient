package TinderClient

import (
	"github.com/parnurzeal/gorequest"
)

type HttpRequest interface {
	Get(url string) HttpRequest
	Post(url string) HttpRequest
	Set(name, value string) HttpRequest
	Send(value string) HttpRequest
	End() (string, []error)
}

type Request struct {
	request *gorequest.SuperAgent
}

func NewRequest() *Request {
	return &Request{request: gorequest.New()}
}

func (r *Request) Get(url string) HttpRequest {
	r.request.Get(url)
	return r
}

func (r *Request) Post(url string) HttpRequest {
	r.request.Post(url)
	return r
}

func (r *Request) Set(name, value string) HttpRequest {
	r.request.Set(name, value)
	return r
}

func (r *Request) Send(value string) HttpRequest {
	r.request.Send(value)
	return r
}

func (r *Request) End() (string, []error) {
	_, body, errs := r.request.End()
	return body, errs
}

type Config struct {
	token      string
	apiBaseUrl string
}

type TRequest struct {
	requester HttpRequest
	token     string
}

func (t *TRequest) Get(url string) (string, []error) {
	return t.requester.Get("https://api.gotinder.com/v2/"+url).
		Set("Authorization", "Token token=\""+t.token+"\"").
		Set("User-Agent", "TinderApp/1.9.1 (iPhone; iOS 14.2; Scale/3.00)").
		Set("X-Auth-Token", t.token).
		End()
}
func (t *TRequest) Post(url string, data string) (string, []error) {
	return t.requester.Post("https://api.gotinder.com/v2/"+url).
		Send(data).
		Set("Authorization", "Token token=\""+t.token+"\"").
		Set("User-Agent", "Tinder/6.9.1 (iPhone; iOS 10.2; Scale/2.00)").
		Set("X-Auth-Token", t.token).
		End()
}

func (t *TRequest) SetHeader(key string, value string) {
	t.requester.Set(key, value)
}

func NewTRequest(token string) *TRequest {
	return &TRequest{
		requester: NewRequest(),
		token:     token,
	}
}
