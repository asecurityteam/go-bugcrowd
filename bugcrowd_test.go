package bugcrowd

import "testing"

func TestNewClient(t *testing.T) {

	auth := BasicAuth{
		Username: "dgd",
		Password: "dgd",
	}

	NewClient(auth, nil)
}
