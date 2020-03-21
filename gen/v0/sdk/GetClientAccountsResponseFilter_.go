// generated by metactl sdk gen 
package sdk

const (
	GetClientAccountsResponseFilterName = "GetClientAccountsResponseFilter"
)

type GetClientAccountsResponseFilter struct {
    And []GetClientAccountsResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientAccounts *ClientAccountListFilter `json:"clientAccounts,omitempty" yaml:"clientAccounts,omitempty"`
    Meta *CollectionMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []GetClientAccountsResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetClientAccountsResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}