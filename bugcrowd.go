package bugcrowd

import (
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.bugcrowd.com"
)

// Client represents the basic struct for the Bugcrowd client
type Client struct {
	Auth    Authentication
	BaseURL *url.URL
	Bounty  *BountyService
	Http    *http.Client
}

// NewClient generates a new client to make outgoing calls to Bugcrowd
func NewClient() (*Client, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		Auth:    Authentication{},
		BaseURL: parsedBaseURL,
		Http:    http.DefaultClient,
	}
	c.Bounty = &BountyService{Client: c}

	return c, nil
}
