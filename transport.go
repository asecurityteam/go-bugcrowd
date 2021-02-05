package bugcrowd

import "net/http"

// Transport placeholder
type Transport struct {
	Authorization     BasicAuth
	OriginalTransport http.RoundTripper
}

// RoundTrip placeholder
func (st *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(st.Authorization.Username, st.Authorization.Password)

	return st.OriginalTransport.RoundTrip(r)
}

// NewTransport placeholder
func NewTransport(auth BasicAuth, rt http.RoundTripper) *http.Client {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return &http.Client{
		Transport: &Transport{
			Authorization:     auth,
			OriginalTransport: rt,
		},
	}
}
