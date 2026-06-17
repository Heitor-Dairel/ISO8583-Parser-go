package custom

import (
	"time"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"
)

type CustomIso struct {
	FileName types.File
	Cycle    types.File
	Date     types.File
	Lines    types.Lines
	Size     types.Size
	Elapsed  time.Duration
}
