// generated by metactl sdk gen 
package sdk

const (
	ClientAccountRelationsSelectName = "ClientAccountRelationsSelect"
)

type ClientAccountRelationsSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    OwnsServiceAccounts *ServiceAccountsCollectionSelect `json:"ownsServiceAccounts,omitempty" yaml:"ownsServiceAccounts,omitempty"`
}