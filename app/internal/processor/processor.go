package processor

import (
	"context"
	"fmt"
	"log"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/filepathiso"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/iso"
	"golang.org/x/sync/errgroup"
)

func ParseISO85831993(files ...string) {

	iso.StartApp()

	var err error
	var errg *errgroup.Group

	errg, _ = errgroup.WithContext(context.Background())

	for i := 0; i < len(files); i += 2 {

		errg.Go(func() error {

			var prs *iso.ISO8583

			prs = iso.NewParse()

			return prs.ParseIso(files[i], files[i+1])

		})

	}

	if err = errg.Wait(); err != nil {
		log.Println(err)
	}

	if len(filepathiso.FilesNotFound) != 0 {
		fmt.Println(filepathiso.FilesNotFound)
	}

}
