// generated by metactl sdk gen 
package sdk

const (
	ServiceListFilterName = "ServiceListFilter"
)

type ServiceListFilter struct {
    Every *ServiceFilter `json:"every,omitempty" yaml:"every,omitempty"`
    None *ServiceFilter `json:"none,omitempty" yaml:"none,omitempty"`
    Some *ServiceFilter `json:"some,omitempty" yaml:"some,omitempty"`
}