// generated by metactl sdk gen 
package sdk

const (
	PutSocialAccountsRequestName = "PutSocialAccountsRequest"
)

type PutSocialAccountsRequest struct {
    Auth *Auth `json:"auth,omitempty" yaml:"auth,omitempty"`
    Meta *RequestMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *PutMode `json:"mode,omitempty" yaml:"mode,omitempty"`
    Select *PutSocialAccountsResponseSelect `json:"select,omitempty" yaml:"select,omitempty"`
    ServiceFilter *ServiceFilter `json:"serviceFilter,omitempty" yaml:"serviceFilter,omitempty"`
    SocialAccounts []SocialAccount `json:"socialAccounts,omitempty" yaml:"socialAccounts,omitempty"`
}