// generated by metactl sdk gen 
package sdk

const (
	IdListFilterName = "IdListFilter"
)

type IdListFilter struct {
    Every *IdFilter `json:"every,omitempty" yaml:"every,omitempty"`
    None *IdFilter `json:"none,omitempty" yaml:"none,omitempty"`
    Some *IdFilter `json:"some,omitempty" yaml:"some,omitempty"`
}