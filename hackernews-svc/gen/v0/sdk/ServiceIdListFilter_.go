// generated by metactl sdk gen 
package sdk

const (
	ServiceIdListFilterName = "ServiceIdListFilter"
)

type ServiceIdListFilter struct {
    Every *ServiceIdFilter `json:"every,omitempty" yaml:"every,omitempty"`
    None *ServiceIdFilter `json:"none,omitempty" yaml:"none,omitempty"`
    Some *ServiceIdFilter `json:"some,omitempty" yaml:"some,omitempty"`
}