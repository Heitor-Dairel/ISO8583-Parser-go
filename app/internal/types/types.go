package types

import "github.com/iancoleman/orderedmap"

type (
	Path          = string
	File          = string
	Data          = []byte
	Size          = string
	HeaderCsv     = []string
	OuterDataCsv  = [][]string
	NestedDataCsv = []string
	Parse         = []*orderedmap.OrderedMap
	Lines         = string
)
