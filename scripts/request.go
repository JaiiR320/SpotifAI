package scripts

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Method string
	URL    string
	Header map[string]string
	Query  map[string]string
	Object *interface{}
}

func (r *Request) Do() error {
	client := &http.Client{}

	// build query
	query := url.Values{}
	for key, value := range r.Query {
		query.Set(key, value)
	}

	// build request
	req, err := http.NewRequest(r.Method, r.URL, strings.NewReader(query.Encode()))
	if err != nil {
		return err
	}

	// set headers to request
	for key, value := range r.Header {
		req.Header.Set(key, value)
	}

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// unmarshal response body to object
	err = json.Unmarshal(body, r.Object)
	if err != nil {
		return err
	}
	return nil
}

func newRequest(url string) *Request {
	req := &Request{Header: make(map[string]string), Query: make(map[string]string)}
	req.URL = url
	return req
}

func Get(url string) *Request {
	req := newRequest(url)
	req.Method = "GET"
	return req
}

func Post(url string) *Request {
	req := newRequest(url)
	req.Method = "POST"
	return req
}

func (req *Request) WithHeader(key, value string) *Request {
	req.Header[key] = value
	return req
}

func (req *Request) WithQuery(key, value string) *Request {
	req.Query[key] = value
	return req
}

func (req *Request) WithObject(object interface{}) *Request {
	req.Object = &object
	return req
}

// type JSON map[string]string
//
// header := JSON{
// 	"Content-Type":  "application/x-www-form-urlencoded",
// 	"Authorization": "Basic " + idAndSecret,
// }

// query := JSON{
// 	"grant_type":   "authorization_code",
// 	"code":         c.QueryParam("code"),
// 	"redirect_uri": "http://localhost:3000/callback",
// }
