// generated by metactl sdk gen
package sdk

const (
	WhateversCollectionSelectName = "WhateversCollectionSelect"
)

type WhateversCollectionSelect struct {
	All        *bool             `json:"all,omitempty" yaml:"all,omitempty"`
	Count      *bool             `json:"count,omitempty" yaml:"count,omitempty"`
	Errors     *ErrorSelect      `json:"errors,omitempty" yaml:"errors,omitempty"`
	Pagination *PaginationSelect `json:"pagination,omitempty" yaml:"pagination,omitempty"`
	Warnings   *WarningSelect    `json:"warnings,omitempty" yaml:"warnings,omitempty"`
	Whatevers  *WhateverSelect   `json:"whatevers,omitempty" yaml:"whatevers,omitempty"`
}
