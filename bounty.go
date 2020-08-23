package bugcrowd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	getBountiesEndpoint    = "/bounties"
	retrieveBountyEndpoint = "/bounties"
)

// BountyAPI is the interface used for mocking of Bounty API calls
type BountyAPI interface {
	GetBounties() ([]Bounty, error)
	RetrieveBounty(uuid string) (RetrieveBountyResponse, error)
}

type BountyService struct {
	Client *Client
}

// GetBountiesResponse is the wrapper object returned by Bugcrowd in its GetBounty response
type GetBountiesResponse struct {
	Bounties []Bounty `json:"bounties,omitempty"`
}

type RetrieveBountyResponse struct {
	Bounty Bounty `json:"bounty,omitempty"`
}

type Bounty struct {
	UUID                    string             `json:"uuid,omitempty"`
	BountyType              string             `json:"bountytype,omitempty"`
	Code                    string             `json:"code,omitempty"`
	CustomFieldLabels       []CustomFieldLabel `json:"custom_field_labels,omitempty"`
	DescriptionMarkdown     string             `json:"description_markdown,omitempty"`
	Demo                    bool               `json:"demo,omitempty"`
	EndsAt                  time.Time          `json:"ends_at,omitempty"`
	HighRewards             int                `json:"high_rewards,omitempty"`
	LowRewards              int                `json:"low_rewards,omitempty"`
	Participation           string             `json:"participation,omitempty"`
	PointsOnly              bool               `json:"points_only,omitempty"`
	StartsAt                time.Time          `json:"starts_at,omitempty"`
	TargetsOverviewMarkdown string             `json:"targets_overview_markdown,omitempty"`
	Tagline                 string             `json:"tagline,omitempty"`
	TotalPrizePool          string             `json:"total_prize_pool,omitempty"`
	RemainingPrizePool      string             `json:"remaining_prize_pool,omitempty"`
	Trial                   bool               `json:"trial,omitempty"`
	Status                  string             `json:"status,omitempty"`
	ServiceLevel            string             `json:"service_level,omitempty"`
	Organization            Organization       `json:"organization,omitempty"`
}

// CustomFieldLabel represents any custom fields put into a bounty
type CustomFieldLabel struct {
	FieldID   string `json:"field_id,omitempty"`
	FieldName string `json:"field_name,omitempty"`
}

// Organization represents the organization a given bounty belongs.
type Organization struct {
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetBounties retrieves all bounty information from Bugcrowd that the you have access
func (b *BountyService) GetBounties() (GetBountiesResponse, error) {
	u, _ := url.Parse(b.Client.BaseURL.String())
	u.Path = path.Join(u.Path, getBountiesEndpoint)

	req, _ := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	req.Header.Set("Accept", "application/vnd.bugcrowd+json")
	req.SetBasicAuth(b.Client.token.Username, b.Client.token.Password)

	resp, err := b.Client.http.Do(req)
	if err != nil {
		return GetBountiesResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return GetBountiesResponse{}, fmt.Errorf("BugCrowd returned non 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetBountiesResponse{}, err
	}

	var bounties GetBountiesResponse
	json.Unmarshal(body, &bounties)

	return bounties, nil
}

// RetrieveBounty retrieves bounty with the given UUID
// If bounty with given ID is not found, an empty response will be returned with a nil error
func (b *BountyService) RetrieveBounty(uuid string) (RetrieveBountyResponse, error) {
	u, _ := url.Parse(b.Client.BaseURL.String())
	u.Path = path.Join(u.Path, retrieveBountyEndpoint, uuid)

	req, _ := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	req.Header.Set("Accept", "application/vnd.bugcrowd+json")
	req.SetBasicAuth(b.Client.token.Username, b.Client.token.Password)

	resp, err := b.Client.http.Do(req)
	if err != nil {
		return RetrieveBountyResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return RetrieveBountyResponse{}, fmt.Errorf("BugCrowd returned non 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RetrieveBountyResponse{}, err
	}

	var bounty RetrieveBountyResponse
	json.Unmarshal(body, &bounty)

	return bounty, nil
}
