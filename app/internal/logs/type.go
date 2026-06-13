package logs

import (
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"

	"github.com/iancoleman/orderedmap"
)

type Log struct {
	FileDataParse []*orderedmap.OrderedMap
	PathTxt       types.Path
	PathCsv       types.Path
}
