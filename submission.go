package bugcrowd

import (
	"context"
	"net/http"
)

const (
	commonsubmissionsEndpoint         = "/submissions"
	fetchAndUpdateSubmissionsEndpoint = "/submissions/%s"
)

// SubmissionAPI is the interface used for mocking of Bounty API calls
type SubmissionAPI interface {
	GetBounties(ctx context.Context, requestConfig *GetBountiesOptions) (*http.Response, *GetBountiesResponse, error)
	RetrieveBounty(ctx context.Context, uuid string) (*http.Response, *RetrieveBountyResponse, error)
}
