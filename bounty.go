package bugcrowd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

const (
	getBountiesEndpoint = "/bounties"
)

type BugcrowdBountyAPI interface {
	GetBounties(ctx context.Context) ([]Bounty, error)
}

type BountyService struct {
	Client *Client
}

type BountyResponse struct {
	Bounties []Bounty `json:"bounties,omitempty"`
}

type Bounty struct {
	UUID              string             `json:"uuid,omitempty"`
	BountyType        string             `json:"bountytype,omitempty"`
	Code              string             `json:"code,omitempty"`
	CustomFieldLabels []CustomFieldLabel `json:"custom_field_labels,omitempty"`
	Status            string             `json:"status,omitempty"`
}
type CustomFieldLabel struct {
	FieldID   string `json:"field_id,omitempty"`
	FieldName string `json:"field_name,omitempty"`
}

// GetBounties retrieves all bounty information from Bugcrowd that the you have access
func (b *BountyService) GetBounties() (BountyResponse, error) {
	u := b.Client.BaseURL
	u.Path = path.Join(u.Path, getBountiesEndpoint)
	req, _ := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	req.Header.Set("Accept", "application/vnd.bugcrowd+json")
	req.SetBasicAuth(b.Client.token.Username, b.Client.token.Password)

	resp, err := b.Client.http.Do(req)
	if err != nil {
		return BountyResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return BountyResponse{}, fmt.Errorf("BugCrowd returned non 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return BountyResponse{}, err
	}

	var bounties BountyResponse
	json.Unmarshal(body, &bounties)

	return bounties, nil
}
