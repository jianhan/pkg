package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ResponseData struct {
	Code    uint                   `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewResponseData(code uint, message string, data map[string]interface{}) *ResponseData {
	return &ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (d *ResponseData) Error() string {
	return fmt.Sprintf("[%d] %s", d.Code, d.Message)
}

// SendJSONResponse simpley send json response.
func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if _, err = w.Write(body); err != nil {

	}
}

// Do behaves as https://golang.org/pkg/net/http/#Client.Do with the exception that
// the Response.Body does not need to be closed. This function should generally
// only be used when it is already known that the response body will be
// relatively small, as it will be completely read into memory.
func Do(client *http.Client, req *http.Request) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bb := &bytes.Buffer{}
	n, err := io.Copy(bb, resp.Body)
	if err != nil {
		return nil, err
	}
	resp.ContentLength = n
	resp.Body = ioutil.NopCloser(bb)
	return resp, nil
}

// Get behaves as https://golang.org/pkg/net/http/#Get but uses simplehttp.Do() to
// make the request.
func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return Do(http.DefaultClient, req)
}

// Head hehaves as https://golang.org/pkg/net/http/#Head but uses simplehttp.Do() to
// make the request.
func Head(url string) (*http.Response, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return Do(http.DefaultClient, req)
}

// Post behaves as https://golang.org/pkg/net/http/#Post but uses simplehttp.Do() to
// make the request.
func Post(url string, bodyType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	return Do(http.DefaultClient, req)
}

// PostForm behaves as https://golang.org/pkg/net/http/#PostForm but uses simplehttp.Do()
// to make the request.
func PostForm(url string, data url.Values) (*http.Response, error) {
	return Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}
