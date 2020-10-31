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
	BaseURL *url.URL
	Bounty  BountyAPI

	http *http.Client
}

// BasicAuth forms basic auth to be passed in client creation to Auth to Bugcrowd
type BasicAuth struct {
	Username string
	Password string
}

// NewClient generates a new client to make outgoing calls to Bugcrowd
func NewClient(auth BasicAuth) (*Client, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		BaseURL: parsedBaseURL,
		http:    NewTransport(auth),
	}
	c.Bounty = &BountyService{client: c}

	return c, nil
}
