// generated by metactl sdk gen
package sdk

const (
	GetSocialAccountsResponseName = "GetSocialAccountsResponse"
)

type GetSocialAccountsResponse struct {
	Count          *int32          `json:"count,omitempty" yaml:"count,omitempty"`
	Errors         []Error         `json:"errors,omitempty" yaml:"errors,omitempty"`
	Pagination     *Pagination     `json:"pagination,omitempty" yaml:"pagination,omitempty"`
	SocialAccounts []SocialAccount `json:"socialAccounts,omitempty" yaml:"socialAccounts,omitempty"`
	Warnings       []Warning       `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}
