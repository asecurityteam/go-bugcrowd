package bugcrowd

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	commonsubmissionsEndpoint         = "/submissions"
	getSubmissionsEndpoint            = "/bounties/%s/submissions"
	fetchAndUpdateSubmissionsEndpoint = "/submissions/%s"
)

// SubmissionAPI is the interface used for mocking of Bounty API calls
type SubmissionAPI interface {
	GetSubmissions(ctx context.Context, uuid string, requestOptions *GetSubmissionsOptions) (*http.Response, *GetSubmissionsResponse, error)
}

// SubmissionService represents the Submission Service struct itself and all required objects
type SubmissionService struct {
	client *Client
}

// GetSubmissionsResponse is the wrapper object returned by Bugcrowd in its GetBounty response
type GetSubmissionsResponse struct {
	Submissions []*Submission `json:"submissions,omitempty"`
}

// GetSubmissionsOptions represents the URL options available to the GetSubmissions endpoint
type GetSubmissionsOptions struct {
	Search     string `url:"search,omitempty"`
	Assignment string `url:"assignment,omitempty"`
	Duplicate  string `url:"duplicate,omitempty"`
	Severity   string `url:"severity,omitempty"`
	Target     string `url:"target,omitempty"`
	Points     string `url:"points,omitempty"`
	Payments   string `url:"payments,omitempty"`
	Researcher string `url:"researcher,omitempty"`
	Source     string `url:"source,omitempty"`
	TargetType string `url:"target_type,omitempty"`
	BlockedBy  string `url:"blocked_by,omitempty"`
	Retest     string `url:"retest,omitempty"`
	Substate   string `url:"substate,omitempty"`
	Submitted  string `url:"submitted,omitempty"`
	VRT        string `url:"vrt,omitempty"`
	Filter     string `url:"filter,omitempty"`
	Sort       string `url:"sort,omitempty"`
	Offset     string `url:"offset,omitempty"`
	Limit      string `url:"limit,omitempty"`
}

// Submission represents the information provided about a Bugcrowd Bounty
type Submission struct {
	BountyCode                     *string           `json:"bounty_code,omitempty"`
	BugURL                         *string           `json:"bug_url,omitempty"`
	Caption                        *string           `json:"caption,omitempty"`
	CustomFields                   map[string]string `json:"custom_fields,omitempty"`
	CVSSString                     *CVSSObject       `json:"cvss_string,omitempty"`
	DescriptionMarkdown            *string           `json:"description_markdown,omitempty"`
	ExtraInfoMarkdown              *string           `json:"extra_info_markdown,omitempty"`
	FileAttachmentsCount           *int              `json:"file_attachments_count,omitempty"`
	HTTPRequest                    *string           `json:"http_request,omitempty"`
	Identity                       *Identity         `json:"identity,omitempty"`
	Priority                       *int              `json:"priority,omitempty"`
	RemediationAdviceMarkdown      *string           `json:"remediation_advice_markdown,omitempty"`
	ReferenceNumber                *string           `json:"reference_number,omitempty"`
	SubmittedAt                    *time.Time        `json:"submitted_at,omitempty"`
	Source                         *string           `json:"source,omitempty"`
	Substate                       *string           `json:"substate,omitempty"`
	RealSubstate                   *string           `json:"real_substate,omitempty"`
	Title                          *string           `json:"title,omitempty"`
	VRTID                          *string           `json:"vrt_id,omitempty"`
	VRTVersion                     *string           `json:"vrt_version,omitempty"`
	VulnerabilityReferenceMarkdown *string           `json:"vulnerability_references_markdown,omitempty"`
	UUID                           *string           `json:"uuid,omitempty"`
	Bounty                         *Bounty           `json:"bounty,omitempty"`
	DuplicateOf                    *Submission       `json:"duplicate_of,omitempty"`
	Duplicate                      *bool             `json:"duplicate,omitempty"`
	Assignee                       *TrackerUser      `json:"assignee,omitempty"`
	User                           *User             `json:"user,omitempty"`
	MonetaryRewards                *[]MonetaryReward `json:"monetary_rewards,omitempty"`
	Target                         *Target           `json:"target,omitempty"`
}

// CVSSObject represents a CVSSObject in Bugcrowd
type CVSSObject struct {
	Version    *string `json:"version,omitempty"`
	Score      *string `json:"score,omitempty"`
	CVSSString *string `json:"cvss_string,omitempty"`
}

// MonetaryReward test
type MonetaryReward struct {
	Amount *string `json:"amount,omitempty"`
}

// Target test
type Target struct {
	Name             *string `json:"name,omitempty"`
	BusinessPriority *string `json:"business_priority,omitempty"`
}

// Identity test
type Identity struct {
	UUID *string `json:"uuid,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// GetSubmissions retrieves all bounty information from Bugcrowd that the you have access
func (s *SubmissionService) GetSubmissions(ctx context.Context, uuid string, requestOptions *GetSubmissionsOptions) (*http.Response, *GetSubmissionsResponse, error) {
	endPath := fmt.Sprintf(getSubmissionsEndpoint, uuid)

	u, err := buildURL(endPath, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)

	submissions := new(GetSubmissionsResponse)
	resp, err := s.client.DoWithDefault(ctx, req, submissions)
	if err != nil {
		return resp, &GetSubmissionsResponse{}, err
	}

	return resp, submissions, nil
}
