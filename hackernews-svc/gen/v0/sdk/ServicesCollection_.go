// generated by metactl sdk gen 
package sdk

const (
	ServicesCollectionName = "ServicesCollection"
)

type ServicesCollection struct {
    Meta *CollectionMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
    Services []Service `json:"services,omitempty" yaml:"services,omitempty"`
}