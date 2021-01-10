package bugcrowd

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	commonBountiesEndpoint = "/bounties"
)

// BountyAPI is the interface used for mocking of Bounty API calls
type BountyAPI interface {
	GetBounties(ctx context.Context, requestConfig GetBountiesRequestConfig) (*http.Response, *GetBountiesResponse, error)
	RetrieveBounty(ctx context.Context, uuid string) (*http.Response, *RetrieveBountyResponse, error)
}

// BountyService test
type BountyService struct {
	client *Client
}

// GetBountiesRequestConfig test
type GetBountiesRequestConfig struct {
	QueryParams GetBountiesRequestQueryParams
}

// GetBountiesRequestQueryParams test
type GetBountiesRequestQueryParams struct {
	Limit  string
	Offset string
}

// GetBountiesResponse is the wrapper object returned by Bugcrowd in its GetBounty response
type GetBountiesResponse struct {
	Bounties []*Bounty `json:"bounties,omitempty"`
}

// RetrieveBountyResponse test
type RetrieveBountyResponse struct {
	Bounty *Bounty `json:"bounty,omitempty"`
}

// TODO : add stringify

// Bounty test
type Bounty struct {
	UUID                    *string             `json:"uuid,omitempty"`
	BountyType              *string             `json:"bountytype,omitempty"`
	Code                    *string             `json:"code,omitempty"`
	CustomFieldLabels       *[]CustomFieldLabel `json:"custom_field_labels,omitempty"`
	DescriptionMarkdown     *string             `json:"description_markdown,omitempty"`
	Demo                    *bool               `json:"demo,omitempty"`
	EndsAt                  *time.Time          `json:"ends_at,omitempty"`
	HighRewards             *int                `json:"high_rewards,omitempty"`
	LowRewards              *int                `json:"low_rewards,omitempty"`
	Participation           *string             `json:"participation,omitempty"`
	PointsOnly              *bool               `json:"points_only,omitempty"`
	StartsAt                *time.Time          `json:"starts_at,omitempty"`
	TargetsOverviewMarkdown *string             `json:"targets_overview_markdown,omitempty"`
	Tagline                 *string             `json:"tagline,omitempty"`
	TotalPrizePool          *string             `json:"total_prize_pool,omitempty"`
	RemainingPrizePool      *string             `json:"remaining_prize_pool,omitempty"`
	Trial                   *bool               `json:"trial,omitempty"`
	Status                  *string             `json:"status,omitempty"`
	ServiceLevel            *string             `json:"service_level,omitempty"`
	Organization            *Organization       `json:"organization,omitempty"`
}

// Organization represents the organization a given bounty belongs.
type Organization struct {
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetBounties retrieves all bounty information from Bugcrowd that the you have access
func (b *BountyService) GetBounties(ctx context.Context, requestConfig GetBountiesRequestConfig) (*http.Response, *GetBountiesResponse, error) {
	u, _ := url.Parse(b.client.BaseURL.String())
	u.Path = path.Join(u.Path, commonBountiesEndpoint)

	req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", bugcrowdJSONAccept)

	q := req.URL.Query()
	q.Add("limit", requestConfig.QueryParams.Limit)
	q.Add("offset", requestConfig.QueryParams.Offset)
	req.URL.RawQuery = q.Encode()

	bounties := new(GetBountiesResponse)
	resp, err := b.client.Do(ctx, req, bounties)
	if err != nil {
		return resp, &GetBountiesResponse{}, err
	}

	return resp, bounties, nil
}

// RetrieveBounty retrieves bounty with the given UUID
// If bounty with given ID is not found, an empty response will be returned with a nil error
func (b *BountyService) RetrieveBounty(ctx context.Context, uuid string) (*http.Response, *RetrieveBountyResponse, error) {
	u, _ := url.Parse(b.client.BaseURL.String())
	u.Path = path.Join(u.Path, commonBountiesEndpoint, uuid)

	req, _ := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", bugcrowdJSONAccept)

	bounty := new(RetrieveBountyResponse)
	resp, err := b.client.Do(ctx, req, bounty)
	if err != nil {
		return resp, &RetrieveBountyResponse{}, err
	}

	return resp, bounty, nil
}
