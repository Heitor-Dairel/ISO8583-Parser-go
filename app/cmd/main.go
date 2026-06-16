package main

import (
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/iso"
)

func main() {

	iso.NewParse(`C:\Users\heitor.tavares\OneDrive - TRIVALE ADMINISTRACAO LTDA\Operação Processadora - Arquivos CSU`)

	iso.ParseISO85831993("26/05/2025", "CIC1", "26/05/2025", "CIC2", "26/05/2025", "CIC3")

}
