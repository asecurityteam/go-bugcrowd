package bugcrowd

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.bugcrowd.com"

	bugcrowdJSONAccept = "application/vnd.bugcrowd+json"
)

// Client represents the basic struct for the Bugcrowd client
type Client struct {
	BaseURL          *url.URL
	Bounty           BountyAPI
	CustomFieldLabel CustomFieldLabelAPI

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
	c.CustomFieldLabel = &CustomFieldLabelService{client: c}

	return c, nil
}

// Do test
func (c *Client) Do(ctx context.Context, r *http.Request, b interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("must pass a non-nil context")
	}
	r = r.WithContext(ctx)

	resp, err := c.http.Do(r)
	if err != nil {
		// TODO : check if context error ocurred
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, err
	}

	// check if b is null as interface as nullable when passed
	if b != nil {
		if w, ok := b.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(b)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}
