// generated by metactl sdk gen 
package sdk

const (
	SocialAccountsCollectionName = "SocialAccountsCollection"
)

type SocialAccountsCollection struct {
    Meta *CollectionMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
    SocialAccounts []SocialAccount `json:"socialAccounts,omitempty" yaml:"socialAccounts,omitempty"`
}