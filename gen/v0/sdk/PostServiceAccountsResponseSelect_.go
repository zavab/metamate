// generated by metactl sdk gen 
package sdk

const (
	PostServiceAccountsResponseSelectName = "PostServiceAccountsResponseSelect"
)

type PostServiceAccountsResponseSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    Meta *ResponseMetaSelect `json:"meta,omitempty" yaml:"meta,omitempty"`
    ServiceAccounts *ServiceAccountSelect `json:"serviceAccounts,omitempty" yaml:"serviceAccounts,omitempty"`
}