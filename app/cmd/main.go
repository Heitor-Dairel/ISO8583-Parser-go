package main

import (
	"log"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/iso"
)

func main() {

	var (
		prs *iso.Mastercard
		err error
	)

	prs, err = iso.NewParse()
	if err != nil {
		log.Fatal(err)
		return
	}

	prs.ParseIso("26/05/2025", "CIC1")
	prs.ParseIso("26/05/2025", "CIC2")
	prs.ParseIso("26/05/2025", "CIC3")

}
