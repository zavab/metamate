// generated by metactl sdk gen 
package sdk

const (
	GetSocialAccountsResponseSelectName = "GetSocialAccountsResponseSelect"
)

type GetSocialAccountsResponseSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    Count *bool `json:"count,omitempty" yaml:"count,omitempty"`
    Errors *ErrorSelect `json:"errors,omitempty" yaml:"errors,omitempty"`
    Pagination *PaginationSelect `json:"pagination,omitempty" yaml:"pagination,omitempty"`
    SocialAccounts *SocialAccountSelect `json:"socialAccounts,omitempty" yaml:"socialAccounts,omitempty"`
    Warnings *WarningSelect `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}