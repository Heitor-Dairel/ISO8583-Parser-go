package logs

import (
	"log"
	"os"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"
)

var Loggers = log.New(os.Stdout, "\033[91m[ERROR] \033[0m", log.Ldate|log.Ltime)

type LogIso struct {
	FileDate       types.File
	FileCycle      types.File
	FileDataParsed types.Parse
	PathOutputTxt  types.Path
	PathOutputCsv  types.Path
}
