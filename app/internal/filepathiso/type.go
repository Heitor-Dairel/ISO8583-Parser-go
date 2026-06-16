package filepathiso

import (
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"
)

type FilePathIso struct {
	FileData   types.Data
	FilePath   types.File
	FileName   types.File
	FileExt    types.File
	FileSize   types.Size
	PathOutput types.Path
	PathIso    types.Path
}
