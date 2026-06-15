package iso

import (
	"context"
	"errors"
	"log"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/exception"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/filepathiso"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/helpers"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/logs"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"
	"golang.org/x/sync/errgroup"
)

var fpi *filepathiso.FilePathIso = &filepathiso.FilePathIso{}

func NewParse() {
	if err := fpi.GetPathOutput(); err != nil {
		log.Fatal(err)
	}
}

func ParseMastercardIso(files ...string) {

	var err error
	var errg *errgroup.Group

	errg, _ = errgroup.WithContext(context.Background())

	for i := 0; i < len(files); i += 2 {

		errg.Go(func() error {

			return parseIso(files[i], files[i+1])

		})

	}

	if err = errg.Wait(); err != nil {
		log.Println(err)
	}

	if len(filepathiso.FilesNotFound) != 0 {
		log.Println(filepathiso.FilesNotFound)
	}

}

func parseIso(fileDate, fileCycle types.File) error {

	var err error
	var logger logs.LogIso
	var iso ISO8583 = ISO8583{FilePathIso: *fpi, FileDataParsed: make(types.Parse, 0, 6000)}
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

	logger = logs.LogIso{
		FileDataParsed: iso.FileDataParsed,
		PathOutputTxt:  iso.FilePathIso.PathOutputTxt,
		PathOutputCsv:  iso.FilePathIso.PathOutputCsv,
	}

	if err = logger.LoggingIso(); err != nil {
		return err
	}

	return nil

}
