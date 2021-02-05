package bugcrowd

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
	reflect "reflect"

	"github.com/google/go-querystring/query"
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
func NewClient(auth BasicAuth, rt http.RoundTripper) (*Client, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		BaseURL: parsedBaseURL,
		http:    NewTransport(auth, rt),
	}

	c.Bounty = &BountyService{client: c}
	c.CustomFieldLabel = &CustomFieldLabelService{client: c}

	return c, nil
}

// DoWithDefault wraps a call around Do that adds the default Bugcrowd headers
func (c *Client) DoWithDefault(ctx context.Context, r *http.Request, b interface{}) (*http.Response, error) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", bugcrowdJSONAccept)

	return c.Do(ctx, r, b)
}

// Do executes the passed in request and sets default headers to the request
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
		return resp, errors.New("Returned non-200")
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

// buildURL adds the parameters passed in as options. This method was inspired by Google's
// addOptions() in the go-github library (https://github.com/google/go-github)
func buildURL(p string, opts interface{}) (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return u, err
	}
	u.Path = path.Join(u.Path, p)

	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil, nil
	}

	qs, err := query.Values(opts)
	if err != nil {
		return u, err
	}

	u.RawQuery = qs.Encode()
	return u, nil
}
