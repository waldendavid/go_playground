package openlibrary

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpClient struct {
	client *http.Client
	url    string
}

func NewHttpClient(client *http.Client, url string) *HttpClient {
	return &HttpClient{
		client: client,
		url:    url}
}
func (hc *HttpClient) Search(ctx context.Context, sreq SearchRequest) (SearchResponse, error) {
	var url string
	if sreq.Author != "" {
		url = fmt.Sprintf("%s/search.json?author=%s", hc.url, sreq.Author)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return SearchResponse{}, err
	}
	res, err := hc.client.Do(req)
	if err != nil {
		return SearchResponse{}, err
	}
	defer res.Body.Close()
	var sresp SearchResponse
	err = json.NewDecoder(res.Body).Decode(&sresp)
	if err != nil {
		return SearchResponse{}, err
	}
	return sresp, nil
}
