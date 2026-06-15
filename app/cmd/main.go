package main

import "github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/processor"

func main() {

	processor.ParseISO85831993("26/05/2025", "CIC1", "26/05/2025", "CIC2", "26/05/2025", "CIC3")

}
