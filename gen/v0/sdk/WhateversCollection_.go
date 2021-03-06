// generated by metactl sdk gen 
package sdk

const (
	WhateversCollectionName = "WhateversCollection"
)

type WhateversCollection struct {
    Count *int32 `json:"count,omitempty" yaml:"count,omitempty"`
    Errors []Error `json:"errors,omitempty" yaml:"errors,omitempty"`
    Pagination *Pagination `json:"pagination,omitempty" yaml:"pagination,omitempty"`
    Warnings []Warning `json:"warnings,omitempty" yaml:"warnings,omitempty"`
    Whatevers []Whatever `json:"whatevers,omitempty" yaml:"whatevers,omitempty"`
}