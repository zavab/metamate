// generated by metactl sdk gen 
package sdk

const (
	UrlFilterName = "UrlFilter"
)

type UrlFilter struct {
    And []UrlFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []UrlFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []UrlFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Value *StringFilter `json:"value,omitempty" yaml:"value,omitempty"`
}