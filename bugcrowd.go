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
	Bounty  *BountyService

	http  *http.Client
	token BasicAuth
}

// NewClient generates a new client to make outgoing calls to Bugcrowd
func NewClient(auth BasicAuth) (*Client, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		BaseURL: parsedBaseURL,
		http:    http.DefaultClient,
		token:   auth,
	}
	c.Bounty = &BountyService{Client: c}

	return c, nil
}

type BasicAuth struct {
	Username string
	Password string
}
