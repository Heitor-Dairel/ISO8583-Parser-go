package iso

import (
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/filepathiso"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/logs"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"

	"log"
)

func NewParse() (*Mastercard, error) {

	var err error
	var master *Mastercard = &Mastercard{}
	var log *logs.Log = &logs.Log{}
	var fpi *filepathiso.FilePathIso = &filepathiso.FilePathIso{}

	err = fpi.GetPathOutput()

	if err != nil {
		return nil, err
	}

	master.FileLog, master.FilePathIso = log, fpi

	return master, nil

}

func (iso *Mastercard) ParseIso(fileDate, fileCycle types.File) {

	var err error

	err = iso.FilePathIso.GetFile(fileDate, fileCycle)

	if err != nil {
		log.Fatal(err)
	}

	err = iso.fileIsoPayload()

	if err != nil {
		log.Fatal(err)
	}

	iso.FileLog.FileDataParse = iso.FileDataParse
	iso.FileLog.PathTxt = iso.FilePathIso.PathOutputTxt
	iso.FileLog.PathCsv = iso.FilePathIso.PathOutputCsv

	err = iso.FileLog.LoggingIso()

	if err != nil {
		log.Fatal(err)
	}

}
