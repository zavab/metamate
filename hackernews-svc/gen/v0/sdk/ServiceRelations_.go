// generated by metactl sdk gen 
package sdk

const (
	ServiceRelationsName = "ServiceRelations"
)

type ServiceRelations struct {
    UsesServiceAccounts *ServiceAccountsCollection `json:"usesServiceAccounts,omitempty" yaml:"usesServiceAccounts,omitempty"`
}