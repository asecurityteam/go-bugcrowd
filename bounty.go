package bugcrowd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

type BugcrowdBountyAPI interface {
	GetBounties(ctx context.Context) error
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

func (b *BountyService) GetBounties() error {
	endpoint := "/bounties"
	u := b.Client.BaseURL
	u.Path = path.Join(u.Path, endpoint)
	req, _ := http.NewRequest(http.MethodGet, u.String(), http.NoBody)

	req.Header.Set("Accept", "application/vnd.bugcrowd+json")

	resp, err := b.Client.Http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Jira returned non 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
