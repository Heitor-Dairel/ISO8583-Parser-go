package filepathiso

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/exception"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/helpers"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"
)

var FilesNotFound []string = make([]string, 0, 100)

func (fp *FilePathIso) GetFile(fileDate, fileCycle types.File) error {

	var err error

	if err = fp.walkFileIso(fileDate, fileCycle); err != nil {
		return err
	}

	if err = fp.readFileIso(); err != nil {
		return err
	}

	return nil

}

func (fp *FilePathIso) GetPathOutput() error {

	const (
		outputDir         string = "output"
		errorMsgCreateDir string = "Erro ao criar diretório '%s' : %w."
	)

	var err error

	fp.PathOutput = filepath.Join(fp.PathAbs, outputDir)

	if err = fp.removePathOutput(); err != nil {
		return err
	}

	if err = os.MkdirAll(fp.PathOutput, 0755); err != nil {
		return fmt.Errorf(errorMsgCreateDir, fp.PathOutput, err)
	}

	return nil
}

func (fp *FilePathIso) GetPathAbs() error {

	const errorMsgExe string = "Erro ao encontrar o diretório do executável : %w."

	var (
		exe string
		err error
	)

	if exe, err = os.Executable(); err != nil {
		return fmt.Errorf(errorMsgExe, err)
	}

	fp.PathAbs = filepath.Dir(exe)

	return nil

}

func (fp *FilePathIso) PathIsoExist() error {

	const errMsgExistDir = "Diretório '%s' não encontrado"
	var err error

	if _, err = os.Stat(fp.PathIso); err != nil {
		return fmt.Errorf(errMsgExistDir, fp.PathIso)
	}

	return nil

}

func (fp *FilePathIso) walkFileIso(fileDate, fileCycle types.File) error {

	const (
		errorMsgDir          string = "Erro ao percorrer diretório '%s' : %w."
		errorMsgFile         string = "Arquivo não encontrado para data '%s' e ciclo '%s' : %w."
		errorMsgFileNotFound string = "Arquivo não encontrado para data '%s' e ciclo '%s'."
	)

	var (
		date            string = helpers.RestoreOriginalDate(fileDate)
		fileNamePreview string = fmt.Sprintf("CSU_ACQ_MASTER_OUTGOING_%s_%s_", fileCycle, date)
		errWalk         error
		fileCollected   bool
		funcWalkIsoDir  func(path string, d os.DirEntry, err error) error
	)

	funcWalkIsoDir = func(path string, d os.DirEntry, err error) error {

		var (
			name, ext        string
			fileNameComplete string = d.Name()
			isMatch          bool   = !d.IsDir() && !strings.Contains(path, "(1)") && strings.Contains(fileNameComplete, fileNamePreview)
		)

		if err != nil {
			return err
		}

		if isMatch {

			ext = filepath.Ext(fileNameComplete)
			name = strings.TrimSuffix(fileNameComplete, ext)

			fp.FilePath, fp.FileName, fp.FileExt, fp.FileFullName, fileCollected = path, name, ext, fileNameComplete, true

			fp.FileDateTime = helpers.FormatDateTime(name)

			return filepath.SkipAll
		}

		return nil
	}

	if errWalk = filepath.WalkDir(fp.PathIso, funcWalkIsoDir); errWalk != nil {
		return fmt.Errorf(errorMsgDir, fp.PathIso, errWalk)
	}

	if !fileCollected {

		FilesNotFound = append(FilesNotFound, fmt.Sprintf(errorMsgFileNotFound, fileDate, fileCycle))

		return fmt.Errorf(errorMsgFile, fileDate, fileCycle, exception.ErrFileNotFound)
	}

	return nil

}

func (fp *FilePathIso) readFileIso() error {
	const (
		errorMsgRead string = "Erro ao ler arquivo '%s' : %w."
		errorMsgSize string = "Erro ao obter o tamanho do arquivo '%s' : %w."
	)

	var (
		err  error
		info os.FileInfo
	)

	if info, err = os.Stat(fp.FilePath); err != nil {
		return fmt.Errorf(errorMsgSize, fp.FilePath, err)
	}

	fp.FileSize = helpers.FormatSize(info.Size())

	if fp.FileData, err = os.ReadFile(fp.FilePath); err != nil {
		return fmt.Errorf(errorMsgRead, fp.FilePath, err)
	}

	return nil
}

func (fp *FilePathIso) removePathOutput() error {

	const errorMsgRemove string = "Erro ao remover diretório '%s' : %w."

	var err error

	if err = os.RemoveAll(fp.PathOutput); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf(errorMsgRemove, fp.PathOutput, err)
	}

	return nil
}
