package bugcrowd

const (
	customFieldLabelEndpoint = "/bounties/%s/custom_field_labels"
)

// CustomFieldLabelAPI test
type CustomFieldLabelAPI interface {
	GetCustomFieldLabels(uuid string) (GetCustomFieldLabelsResponse, error)
	CreateCustomFieldLabel(uuid string) error
	UpdateCustomFieldLabel(bountyUUID string, customFieldLabelUUID string) error
	DeleteCustomFieldLabel(bountUUID string, customFieldLabelUUID string) error
}

// CustomFieldLabelService test
type CustomFieldLabelService struct {
	Client *Client
}

// CustomFieldLabel represents any custom fields put into a bounty
type CustomFieldLabel struct {
	FieldID   string `json:"field_id,omitempty"`
	FieldName string `json:"field_name,omitempty"`
}

// GetCustomFieldLabelsResponse test
type GetCustomFieldLabelsResponse struct {
	CustomFieldLabels []CustomFieldLabel `json:"custom_field_labels,omitempty"`
}

// GetCustomFieldLabels test
// func (c *CustomFieldLabelService) GetCustomFieldLabels(uuid string) (GetCustomFieldLabelsResponse, error) {
// 	u, _ := url.Parse(c.Client.BaseURL.String())
// 	u.Path = path.Join(u.Path, customFieldLabelEndpoint)

// 	req, _ := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
// 	req.Header.Set("Accept", "application/vnd.bugcrowd+json")
// 	req.SetBasicAuth(c.Client.auth.Username, c.Client.auth.Password)

// 	resp, err := c.Client.http.Do(req)
// 	if err != nil {
// 		return GetCustomFieldLabelsResponse{}, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		return GetCustomFieldLabelsResponse{}, fmt.Errorf("BugCrowd returned non 200: %d", resp.StatusCode)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return GetCustomFieldLabelsResponse{}, err
// 	}

// 	var customFieldLabels GetCustomFieldLabelsResponse
// 	json.Unmarshal(body, &customFieldLabels)

// 	return customFieldLabels, nil
// }
