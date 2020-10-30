package bugcrowd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type errReader struct {
	Error error
}

func (r *errReader) Read(_ []byte) (int, error) {
	return 0, r.Error
}

func (r *errReader) UnmarshalJSON(b []byte) error {
	return errors.New("Unmarshal Error")
}

func TestGetBounties(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRT := NewMockRoundTripper(ctrl)
	u, _ := url.Parse("http://localhost")

	mockBountyService := NewMockBountyAPI(ctrl)

	client := Client{
		BaseURL: u,
		Bounty:  mockBountyService,

		http: &http.Client{Transport: mockRT},
	}

	bountyService := BountyService{
		client: &client,
	}

	mockJSON := Organization{}

	respJSON, _ := json.Marshal(mockJSON)
	respReader := bytes.NewReader(respJSON)

	tests := []struct {
		name        string
		response    *http.Response
		responseErr error
		expectErr   bool
	}{
		{
			name: "success",
			response: &http.Response{
				Body:       ioutil.NopCloser(respReader),
				StatusCode: http.StatusOK,
			},
			responseErr: nil,
			expectErr:   false,
		},
		{
			name: "failure non 200",
			response: &http.Response{
				Body:       ioutil.NopCloser(respReader),
				StatusCode: http.StatusTeapot,
			},
			responseErr: nil,
			expectErr:   true,
		},
		{
			name:        "request error",
			response:    nil,
			responseErr: errors.New("HTTPError"),
			expectErr:   true,
		},
		{
			name: "io read error",
			response: &http.Response{
				Body:       ioutil.NopCloser(&errReader{Error: fmt.Errorf("io read error")}),
				StatusCode: http.StatusOK,
			},
			responseErr: nil,
			expectErr:   true,
		},
		{
			name: "failure unmarshalling",
			response: &http.Response{
				Body:       ioutil.NopCloser(respReader),
				StatusCode: http.StatusCreated,
			},
			responseErr: nil,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRT.EXPECT().RoundTrip(gomock.Any()).Return(tt.response, tt.responseErr)
			// mockBountyService.EXPECT().GetBounties(gomock.Any())
			_, err := bountyService.GetBounties(GetBountiesRequestConfig{})
			if tt.expectErr {
				require.Error(t, err)
				return
			}
			require.Nil(t, err)
		})
	}
}

func TestRetrieveBounty(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRT := NewMockRoundTripper(ctrl)
	u, _ := url.Parse("http://localhost")

	mockBountyService := NewMockBountyAPI(ctrl)

	client := Client{
		BaseURL: u,
		Bounty:  mockBountyService,

		http: &http.Client{Transport: mockRT},
	}

	bountyService := BountyService{
		client: &client,
	}

	mockJSON := Organization{}

	respJSON, _ := json.Marshal(mockJSON)
	respReader := bytes.NewReader(respJSON)

	tests := []struct {
		name        string
		response    *http.Response
		responseErr error
		expectErr   bool
	}{
		{
			name: "success",
			response: &http.Response{
				Body:       ioutil.NopCloser(respReader),
				StatusCode: http.StatusOK,
			},
			responseErr: nil,
			expectErr:   false,
		},
		{
			name: "failure non 200",
			response: &http.Response{
				Body:       ioutil.NopCloser(respReader),
				StatusCode: http.StatusTeapot,
			},
			responseErr: nil,
			expectErr:   true,
		},
		{
			name:        "request error",
			response:    nil,
			responseErr: errors.New("HTTPError"),
			expectErr:   true,
		},
		{
			name: "io read error",
			response: &http.Response{
				Body:       ioutil.NopCloser(&errReader{Error: fmt.Errorf("io read error")}),
				StatusCode: http.StatusOK,
			},
			responseErr: nil,
			expectErr:   true,
		},
		{
			name: "failure unmarshalling",
			response: &http.Response{
				Body:       ioutil.NopCloser(respReader),
				StatusCode: http.StatusCreated,
			},
			responseErr: nil,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRT.EXPECT().RoundTrip(gomock.Any()).Return(tt.response, tt.responseErr)
			// mockBountyService.EXPECT().GetBounties(gomock.Any())
			_, err := bountyService.RetrieveBounty("any")
			if tt.expectErr {
				require.Error(t, err)
				return
			}
			require.Nil(t, err)
		})
	}
}
