package iso

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/custom"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/exception"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/filepathiso"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/helpers"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/logs"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"
	"golang.org/x/sync/errgroup"
)

var filePathIso filepathiso.FilePathIso = filepathiso.FilePathIso{}

func NewParse(pathIso types.Path) {

	filePathIso.PathIso = pathIso
	if err := filePathIso.GetPathOutput(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(types.Title)
}

func ParseISO85831993(files ...string) {

	var err error
	var errg *errgroup.Group

	errg, _ = errgroup.WithContext(context.Background())

	for i := 0; i < len(files); i += 2 {

		errg.Go(func() error {

			return processorIso(files[i], files[i+1])

		})

	}

	if err = errg.Wait(); err != nil {
		log.Println(err)
	}

	if len(filepathiso.FilesNotFound) != 0 {
		log.Println(filepathiso.FilesNotFound)
	}

}

func processorIso(fileDate, fileCycle types.File) error {

	var (
		start     time.Time = time.Now()
		err       error
		logger    logs.LogIso
		iso       ISO8583          = ISO8583{FilePathIso: filePathIso, FileDataParsed: make(types.Parse, 0, 6000)}
		valid     helpers.Validate = helpers.Validate{FileDate: fileDate, FileCycle: fileCycle}
		customIso custom.CustomIso
	)

	if err = valid.ValidateFileInfoIso(); err != nil {
		return err
	}

	err = iso.FilePathIso.GetFile(fileDate, fileCycle)

	if err != nil && errors.Is(err, exception.ErrFileNotFound) {
		return nil
	}

	if err != nil && !errors.Is(err, exception.ErrFileNotFound) {
		return err
	}

	if err = iso.fileIsoPayload(); err != nil {
		return err
	}

	logger = logs.LogIso{
		FileDate:       fileDate,
		FileCycle:      fileCycle,
		FileDataParsed: iso.FileDataParsed,
		PathOutputTxt:  filepath.Join(iso.FilePathIso.PathOutput, fmt.Sprintf("%s.TXT.LOG", iso.FilePathIso.FileName)),
		PathOutputCsv:  filepath.Join(iso.FilePathIso.PathOutput, fmt.Sprintf("%s.CSV", iso.FilePathIso.FileName)),
	}

	if err = logger.LoggingIso(); err != nil {
		return err
	}

	customIso = custom.CustomIso{
		FileName: iso.FilePathIso.FileFullName,
		Cycle:    fileCycle,
		Date:     iso.FilePathIso.FileDateTime,
		Lines:    iso.ParsedLines,
		Size:     iso.FilePathIso.FileSize,
		Elapsed:  time.Since(start),
	}

	customIso.PrintLog()

	return nil

}
