package filepathiso

import (
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"
)

type FilePathIso struct {
	FileData     types.Data
	FilePath     types.File
	FileName     types.File
	FileFullName types.File
	FileExt      types.File
	FileSize     types.Size
	FileDateTime types.File
	PathAbs      types.Path
	PathOutput   types.Path
	PathIso      types.Path
}
