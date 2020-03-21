// generated by metactl sdk gen 
package sdk

const (
	GetServiceAccountsRequestName = "GetServiceAccountsRequest"
)

type GetServiceAccountsRequest struct {
    Auth *Auth `json:"auth,omitempty" yaml:"auth,omitempty"`
    Filter *ServiceAccountFilter `json:"filter,omitempty" yaml:"filter,omitempty"`
    Meta *RequestMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *GetMode `json:"mode,omitempty" yaml:"mode,omitempty"`
    Pages []ServicePage `json:"pages,omitempty" yaml:"pages,omitempty"`
    Relations *GetServiceAccountsRelations `json:"relations,omitempty" yaml:"relations,omitempty"`
    Select *GetServiceAccountsResponseSelect `json:"select,omitempty" yaml:"select,omitempty"`
    ServiceFilter *ServiceFilter `json:"serviceFilter,omitempty" yaml:"serviceFilter,omitempty"`
    Sort *ServiceAccountSort `json:"sort,omitempty" yaml:"sort,omitempty"`
}