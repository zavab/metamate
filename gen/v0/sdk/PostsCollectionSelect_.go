// generated by metactl sdk gen
package sdk

const (
	PostsCollectionSelectName = "PostsCollectionSelect"
)

type PostsCollectionSelect struct {
	All        *bool             `json:"all,omitempty" yaml:"all,omitempty"`
	Count      *bool             `json:"count,omitempty" yaml:"count,omitempty"`
	Errors     *ErrorSelect      `json:"errors,omitempty" yaml:"errors,omitempty"`
	Pagination *PaginationSelect `json:"pagination,omitempty" yaml:"pagination,omitempty"`
	Posts      *PostSelect       `json:"posts,omitempty" yaml:"posts,omitempty"`
	Warnings   *WarningSelect    `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}
