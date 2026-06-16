package logs

import "github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"

type LogIso struct {
	FileDate       types.File
	FileCycle      types.File
	FileDataParsed types.Parse
	PathOutputTxt  types.Path
	PathOutputCsv  types.Path
}
