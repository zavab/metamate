// generated by metactl sdk gen 
package sdk

const (
	PostClientAccountsResponseSelectName = "PostClientAccountsResponseSelect"
)

type PostClientAccountsResponseSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    ClientAccounts *ClientAccountSelect `json:"clientAccounts,omitempty" yaml:"clientAccounts,omitempty"`
    Meta *ResponseMetaSelect `json:"meta,omitempty" yaml:"meta,omitempty"`
}