// generated by metactl sdk gen 
package sdk

const (
	InfoFilterName = "InfoFilter"
)

type InfoFilter struct {
    And []InfoFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Description *TextFilter `json:"description,omitempty" yaml:"description,omitempty"`
    Name *TextFilter `json:"name,omitempty" yaml:"name,omitempty"`
    Not []InfoFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []InfoFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Purpose *TextFilter `json:"purpose,omitempty" yaml:"purpose,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}