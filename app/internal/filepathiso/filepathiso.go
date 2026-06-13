package filepathiso

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/helpers"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"
)

func (valid *validateSearch) validate() error {
	const (
		cycle1        string = "CIC1"
		cycle2        string = "CIC2"
		cycle3        string = "CIC3"
		formatPrev    string = "02/01/2006"
		errorMsgDate  string = `Formato de data inválido: '%s'.`
		errorMsgCycle string = `Ciclo inválido: '%s'. Os ciclos disponíveis são: CIC1, CIC2, CIC3.`
	)

	var err error

	_, err = time.Parse(formatPrev, valid.fileDate)
	if err != nil {
		return fmt.Errorf(errorMsgDate, valid.fileDate)
	}

	switch valid.fileCycle {
	case cycle1, cycle2, cycle3:
		return nil
	}

	return fmt.Errorf(errorMsgCycle, valid.fileCycle)

}

func (fp *FilePathIso) GetFile(fileDate, fileCycle types.File) error {
	var (
		err         error
		validSearch validateSearch = validateSearch{fileDate: fileDate, fileCycle: fileCycle}
	)

	if err = validSearch.validate(); err != nil {
		return err
	}

	if err = fp.walkFileIso(fileDate, fileCycle); err != nil {
		return err
	}

	fp.pathOutputFiles()

	err = fp.readFileIso()

	if err != nil {
		return err
	}

	return nil

}

func (fp *FilePathIso) pathOutputFiles() {

	fp.PathOutputTxt = filepath.Join(fp.PathOutput, fmt.Sprintf("%s.txt.log", fp.FileName))
	fp.PathOutputCsv = filepath.Join(fp.PathOutput, fmt.Sprintf("%s.csv", fp.FileName))

}

func (fp *FilePathIso) walkFileIso(fileDate, fileCycle types.File) error {
	const (
		pathIsoFiles string = `C:\Users\heitor.tavares\OneDrive - TRIVALE ADMINISTRACAO LTDA\Operação Processadora - Arquivos CSU`
		errorMsgDir  string = `Erro ao percorrer diretório '%s' : %w.`
		errorMsgFile string = `Arquivo não encontrado para ciclo '%s' e data %s.`
	)

	var (
		date           string = helpers.RestoreOriginalDate(fileDate)
		err            error
		name, ext      string
		isFileCollect  bool
		funcWalkIsoDir func(path string, d os.DirEntry, err error) error
	)
	funcWalkIsoDir = func(path string, d os.DirEntry, err error) error {

		var (
			fileNamePreview  string = fmt.Sprintf("CSU_ACQ_MASTER_OUTGOING_%s_%s_", fileCycle, date)
			fileNameComplete string = d.Name()
			isMatch          bool   = !d.IsDir() && !strings.Contains(path, "(1)") && strings.Contains(fileNameComplete, fileNamePreview)
		)

		if err != nil {
			return err
		}

		if isMatch {

			ext = filepath.Ext(fileNameComplete)
			name = strings.TrimSuffix(fileNameComplete, ext)

			fp.FilePath, fp.FileName, fp.FileExt, isFileCollect = path, name, ext, true

			return filepath.SkipAll
		}

		return nil
	}

	err = filepath.WalkDir(pathIsoFiles, funcWalkIsoDir)

	if err != nil {
		return fmt.Errorf(errorMsgDir, pathIsoFiles, err)
	}

	if !isFileCollect {
		return fmt.Errorf(errorMsgFile, fileCycle, fileDate)
	}

	return nil

}

func (fp *FilePathIso) readFileIso() error {
	const errorMsg string = `Erro ao ler arquivo '%s' : %w.`
	var err error

	fp.FileData, err = os.ReadFile(fp.FilePath)
	if err != nil {
		return fmt.Errorf(errorMsg, fp.FilePath, err)
	}

	return nil
}

func (fp *FilePathIso) GetPathOutput() error {
	const outputDir string = "output"
	var (
		exe string
		err error
	)

	exe, err = os.Executable()
	if err != nil {
		return err
	}

	fp.PathOutput = filepath.Join(filepath.Dir(exe), outputDir)

	err = os.RemoveAll(fp.PathOutput)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	err = os.MkdirAll(fp.PathOutput, 0755)

	if err != nil {
		return err
	}

	return nil
}
