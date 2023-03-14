package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// CreatePostRequest takes a url,headers and a body and create a post request.
func CreatePostRequest(ctx context.Context, url string, data interface{}, headers map[string]string) (*http.Request, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, err
}
