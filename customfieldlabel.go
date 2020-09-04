package bugcrowd

const (
	customFieldEndpoint = "/bounties/%s/custom_field_labels"
)

// CustomFieldLabelAPI test
type CustomFieldLabelAPI interface {
	GetCustomFieldLabels(uuid string)
	CreateCustomFieldLabel(uuid string)
	UpdateCustomFieldLabel(bountyUUID string, customFieldLabelUUID string)
	DeleteCustomFieldLabel(bountUUID string, customFieldLabelUUID string)
}

// CustomFieldLabelService test
type CustomFieldLabelService struct {
	Client *Client
}
