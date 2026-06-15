package iso

import (
	"errors"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/exception"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/filepathiso"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/helpers"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/logs"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"

	"log"
)

var fpi *filepathiso.FilePathIso = &filepathiso.FilePathIso{}

func StartApp() {

	var err error

	if err = fpi.GetPathOutput(); err != nil {
		log.Fatal(err)
	}

}

func NewParse() *ISO8583 {

	return &ISO8583{FilePathIso: *fpi, FileDataParsed: make(types.Parse, 0, 6000)}

}

func (iso *ISO8583) ParseIso(fileDate, fileCycle types.File) error {

	var err error
	var logger *logs.LogIso
	var valid helpers.Validate = helpers.Validate{FileDate: fileDate, FileCycle: fileCycle}

	if err = valid.ValidateFileInfoIso(); err != nil {
		return err
	}

	err = iso.FilePathIso.GetFile(fileDate, fileCycle)

	if err != nil && !errors.Is(err, exception.ErrFileNotFound) {
		return err
	}

	if err != nil && errors.Is(err, exception.ErrFileNotFound) {
		return nil
	}

	if err = iso.fileIsoPayload(); err != nil {
		return err
	}

	logger = &logs.LogIso{
		FileDataParsed: iso.FileDataParsed,
		PathOutputTxt:  iso.FilePathIso.PathOutputTxt,
		PathOutputCsv:  iso.FilePathIso.PathOutputCsv,
	}

	if err = logger.LoggingIso(); err != nil {
		return err
	}

	return nil

}
