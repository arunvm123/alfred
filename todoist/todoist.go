package todoist

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	endpoint = "https://api.todoist.com/rest/v1"
)

type Client struct {
	BaseURL    string
	Token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		BaseURL:    endpoint,
		Token:      token,
		httpClient: http.DefaultClient,
	}
}

func (c *Client) postMethod(path string, body io.Reader, res interface{}) error {
	req, err := http.NewRequest("POST", c.BaseURL+path, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	_, err = c.do(req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) getMethod(path string, res interface{}) error {
	req, err := http.NewRequest("GET", c.BaseURL+path, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	_, err = c.do(req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}
