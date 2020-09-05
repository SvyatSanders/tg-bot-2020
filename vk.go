package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version = "5.122"
	reqURL  = "https://api.vk.com/method/wall.get?"
)

type wallResponse struct {
	Body body `json:"response"`
}

type body struct {
	Items []items `json:"items"`
}

type items struct {
	Text string `json:"text"`
}

func getPostsQuery(groupID string, vkServiceKey string) ([]items, error) {
	u := url.Values{}
	u.Set("count", "3")
	u.Set("offset", "0")
	u.Set("filter", "owner")
	u.Set("owner_id", groupID)
	u.Set("access_token", vkServiceKey)
	u.Set("v", version)

	req := reqURL + u.Encode()
	resp, err := http.Get(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(wallResponse)
	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response.Body.Items, nil
}
