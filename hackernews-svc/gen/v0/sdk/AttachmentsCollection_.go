// generated by metactl sdk gen 
package sdk

const (
	AttachmentsCollectionName = "AttachmentsCollection"
)

type AttachmentsCollection struct {
    Attachments []Attachment `json:"attachments,omitempty" yaml:"attachments,omitempty"`
    Count *int32 `json:"count,omitempty" yaml:"count,omitempty"`
    Errors []Error `json:"errors,omitempty" yaml:"errors,omitempty"`
    Pagination *Pagination `json:"pagination,omitempty" yaml:"pagination,omitempty"`
    Warnings []Warning `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}