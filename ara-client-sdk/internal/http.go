package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func CreatePostRequest(url string, data interface{}, headers map[string]string) (*http.Request, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	/*client := &http.Client{
		Timeout: timeOut,
	}*/
	return req, err
}

func postRequest(url string, data interface{}, headers map[string]string, timeOut time.Duration) (*http.Response, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: timeOut,
	}
	return client.Do(req)

}
