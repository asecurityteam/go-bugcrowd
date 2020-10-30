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

	http *http.Client
	auth BasicAuth
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
		auth:    auth,
	}
	c.Bounty = &BountyService{client: c}

	return c, nil
}

// NewRequest Creates a new http.Request object with the basic headers/auth required to communicate
// with bugcrowd
// func (c *Client) NewRequest(method, url string, payload io.Reader) (*http.Request, error) {
// 	req, err := http.NewRequest(method, url, payload)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Accept", "application/vnd.bugcrowd+json")
// 	req.SetBasicAuth(c.auth.Username, c.auth.Password)

// 	return req, nil
// }
