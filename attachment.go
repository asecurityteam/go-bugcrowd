package bugcrowd

import (
	"context"
	"fmt"
	"net/http"
)

const (
	attachmentsEndpoint = "/submissions/%s/file_attachments"
)

// AttachmentAPI is the interface used for mocking of Bounty API calls
type AttachmentAPI interface {
	ViewSubmissionAttachments(ctx context.Context, uuid string) (*http.Response, *GetSubmissionsResponse, error)
}

// AttachmentService represents the Attachment Service struct itself and all required objects
type AttachmentService struct {
	client *Client
}

// ViewSubmissionAttachmentsResponse represents the response expected from the GetAttachments API
type ViewSubmissionAttachmentsResponse struct {
	FileAttachments []Attachment `json:"file_attachments,omitempty"`
}

// Attachment represents the structure of an attachment in bugcrowd
type Attachment struct {
	FileName    string `json:"file_name,omitempty"`
	FileSize    int    `json:"file_size,omitempty"`
	FileType    string `json:"file_type,omitempty"`
	S3SignedURL string `json:"s3_signed_url,omitempty"`
}

// RetrieveSubmission retrieves all bounty information from Bugcrowd that the you have access
func (s *AttachmentService) RetrieveSubmission(ctx context.Context, uuid string) (*http.Response, *ViewSubmissionAttachmentsResponse, error) {
	endPath := fmt.Sprintf(retrieveAndUpdateSubmissionsEndpoint, uuid)

	u, err := buildURL(endPath, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), http.NoBody)

	attachments := new(ViewSubmissionAttachmentsResponse)
	resp, err := s.client.DoWithDefault(ctx, req, attachments)
	if err != nil {
		return resp, &ViewSubmissionAttachmentsResponse{}, err
	}

	return resp, attachments, nil
}
