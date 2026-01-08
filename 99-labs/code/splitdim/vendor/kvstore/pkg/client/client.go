package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"kvstore/pkg/api"
)

type client struct {
	url string
}

func NewClient(addr string) Client {
	return &client{url: "http://" + addr}
}

// Get returns the value and version stored for the given key, or an error if something goes wrong.
func (c *client) Get(key string) (api.VersionedValue, error) {
	uri, _ := url.Parse(c.url + "/api/get")
	q := uri.Query()
	q.Set("id", key)
	uri.RawQuery = q.Encode()

	r, err := http.Get(uri.String())
	if err != nil {
		return api.VersionedValue{}, err
	}
	if r.StatusCode != http.StatusOK {
		return api.VersionedValue{}, fmt.Errorf("Request failed with status code %d", r.StatusCode)
	}
	value := api.VersionedValue{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return api.VersionedValue{}, fmt.Errorf("Could not read response body")
	}
	err = json.Unmarshal([]byte(body), &value)
	if err != nil {
		return api.VersionedValue{}, fmt.Errorf("Could not unmarshal response body")
	}
	return value, nil
}

// Put tries to insert the given key-value pair with the specified version into the store.
func (c *client) Put(vkv api.VersionedKeyValue) error {
	json, err := json.Marshal(vkv)
	if err != nil {
		return err
	}
	r, err := http.Post(c.url+"/api/put", "application/json", bytes.NewReader(json))
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("Put request returned with error code %v", r.StatusCode)
	}
	return nil
}

// List returns all values stored in the database.
func (c *client) List() ([]api.VersionedKeyValue, error) {
	r, err := http.Get(c.url + "/api/list")
	if err != nil {
		return []api.VersionedKeyValue{}, err
	}
	if r.StatusCode != http.StatusOK {
		return []api.VersionedKeyValue{}, fmt.Errorf("List returned with error code %v", r.StatusCode)
	}
	value := []api.VersionedKeyValue{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return []api.VersionedKeyValue{}, err
	}
	err = json.Unmarshal([]byte(body), &value)
	if err != nil {
		return []api.VersionedKeyValue{}, err
	}
	return value, nil
}

// Reset removes all key-value pairs.
func (c *client) Reset() error {
	r, err := http.Get(c.url + "/api/reset")
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("Reset returned with status code %v", r.StatusCode)
	}
	return nil
}
