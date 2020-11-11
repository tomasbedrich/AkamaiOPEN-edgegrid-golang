package storage

import "context"

// ResponseStatus is returned on Create, Update or Delete operations for all entity types
type ResponseStatus struct {
	ChangeId              string  `json:"changeId,omitempty"`
	Links                 *[]Link `json:"links,omitempty"`
	Message               string  `json:"message,omitempty"`
	PassingValidation     bool    `json:"passingValidation,omitempty"`
	PropagationStatus     string  `json:"propagationStatus,omitempty"`
	PropagationStatusDate string  `json:"propagationStatusDate,omitempty"`
}

// Probably THE most common type
type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type StorageGroupResponse struct {
	Status   *ResponseStatus `json:"status"`
	Resource *StorageGroup   `json:"resource"`
}

// NewStorageGroupResponse instantiates a new StorageGroupResponse structure
func (p *storage) NewStorageGroupResponse(ctx context.Context) *StorageGroupResponse {
	logger := p.Log(ctx)
	logger.Debug("NewStorageGroupResponse")
	resp := &StorageGroupResponse{}
	return resp
}
