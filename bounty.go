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

type Bounty struct {
	UUID       string
	BountyType string
	Code       string
	Status     string
}

// GetBounties retrieves all bounty information from Bugcrowd that the you have access
func (b *BountyService) GetBounties() ([]Bounty, error) {
	u := b.Client.BaseURL
	u.Path = path.Join(u.Path, getBountiesEndpoint)
	req, _ := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	req.Header.Set("Accept", "application/vnd.bugcrowd+json")
	req.SetBasicAuth(b.Client.token.Username, b.Client.token.Password)

	resp, err := b.Client.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []Bounty{}, fmt.Errorf("BugCrowd returned non 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var bounties []Bounty
	json.Unmarshal(body, &bounties)

	return bounties, nil
}
