package main

import (
	"os"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/iso"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/logs"
	"github.com/joho/godotenv"
)

func main() {

	const errorMsg string = "Erro ao carregar arquivo .env."

	err := godotenv.Load()

	if err != nil {

		logs.Loggers.Fatal(errorMsg)

	}

	iso.NewParse(os.Getenv("DIR_PATH"))

	iso.ParseISO85831993("26/05/2025", "CIC2")

}
