package iso

import (
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/filepathiso"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/logs"

	"github.com/iancoleman/orderedmap"
)

type Mastercard struct {
	FilePathIso   *filepathiso.FilePathIso
	FileDataParse []*orderedmap.OrderedMap
	FileLog       *logs.Log
}
