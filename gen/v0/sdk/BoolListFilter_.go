// generated by metactl sdk gen 
package sdk

const (
	BoolListFilterName = "BoolListFilter"
)

type BoolListFilter struct {
    And *BoolFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not *BoolFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or *BoolFilter `json:"or,omitempty" yaml:"or,omitempty"`
}