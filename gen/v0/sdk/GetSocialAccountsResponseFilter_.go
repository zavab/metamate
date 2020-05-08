// generated by metactl sdk gen
package sdk

const (
	GetSocialAccountsResponseFilterName = "GetSocialAccountsResponseFilter"
)

type GetSocialAccountsResponseFilter struct {
	And            []GetSocialAccountsResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
	Count          *Int32Filter                      `json:"count,omitempty" yaml:"count,omitempty"`
	Errors         *ErrorListFilter                  `json:"errors,omitempty" yaml:"errors,omitempty"`
	Not            []GetSocialAccountsResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
	Or             []GetSocialAccountsResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
	Pagination     *PaginationFilter                 `json:"pagination,omitempty" yaml:"pagination,omitempty"`
	Set            *bool                             `json:"set,omitempty" yaml:"set,omitempty"`
	SocialAccounts *SocialAccountListFilter          `json:"socialAccounts,omitempty" yaml:"socialAccounts,omitempty"`
	Warnings       *WarningListFilter                `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}
