// generated by metactl sdk gen 
package sdk

const (
	AttachmentsCollectionName = "AttachmentsCollection"
)

type AttachmentsCollection struct {
    Attachments []Attachment `json:"attachments,omitempty" yaml:"attachments,omitempty"`
    Meta *CollectionMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
}