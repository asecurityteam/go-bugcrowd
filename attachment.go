package bugcrowd

import (
	"context"
	"fmt"
	"net/http"
)

const (
	commentAttachmentsEndpoint = "/submissions/%s/comments/%s/file_attachments"
	attachmentsEndpoint        = "/submissions/%s/file_attachments"
)

// AttachmentAPI is the interface used for mocking of Bounty API calls
type AttachmentAPI interface {
	ViewCommentAttachments(ctx context.Context, uuid string) (*http.Response, *ViewAttachmentsResponse, error)
	ViewSubmissionAttachments(ctx context.Context, uuid string) (*http.Response, *ViewAttachmentsResponse, error)
}

// AttachmentService represents the Attachment Service struct itself and all required objects
type AttachmentService struct {
	client *Client
}

// ViewAttachmentsResponse represents the response expected from the GetAttachments API
type ViewAttachmentsResponse struct {
	FileAttachments []Attachment `json:"file_attachments,omitempty"`
}

// Attachment represents the structure of an attachment in bugcrowd
type Attachment struct {
	FileName    string `json:"file_name,omitempty"`
	FileSize    int    `json:"file_size,omitempty"`
	FileType    string `json:"file_type,omitempty"`
	S3SignedURL string `json:"s3_signed_url,omitempty"`
}

// ViewCommentAttachments retrieves all bounty information from Bugcrowd that the you have access
func (s *AttachmentService) ViewCommentAttachments(ctx context.Context, submissionUUID, commentUUID string) (*http.Response, *ViewAttachmentsResponse, error) {
	endPath := fmt.Sprintf(commentAttachmentsEndpoint, submissionUUID, commentUUID)

	u, err := buildURL(endPath, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, nil, err
	}

	attachments := new(ViewAttachmentsResponse)
	resp, err := s.client.DoWithDefault(ctx, req, attachments)
	if err != nil {
		return resp, &ViewAttachmentsResponse{}, err
	}

	return resp, attachments, nil
}

// ViewSubmissionAttachments retrieves all bounty information from Bugcrowd that the you have access
func (s *AttachmentService) ViewSubmissionAttachments(ctx context.Context, uuid string) (*http.Response, *ViewAttachmentsResponse, error) {
	endPath := fmt.Sprintf(attachmentsEndpoint, uuid)

	u, err := buildURL(endPath, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, nil, err
	}

	attachments := new(ViewAttachmentsResponse)
	resp, err := s.client.DoWithDefault(ctx, req, attachments)
	if err != nil {
		return resp, &ViewAttachmentsResponse{}, err
	}

	return resp, attachments, nil
}
