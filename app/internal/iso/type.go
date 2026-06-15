package iso

import (
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/filepathiso"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"
)

type ISO8583 struct {
	FilePathIso    filepathiso.FilePathIso
	FileDataParsed types.Parse
	ParsedLines    types.Lines
}
