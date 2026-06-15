package main

import (
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/iso"
)

func main() {

	iso.NewParse()

	iso.ParseMastercardIso("26/05/2025", "CIC1", "26/05/2025", "CIC2", "26/05/2025", "CIC3")

}
