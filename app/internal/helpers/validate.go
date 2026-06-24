package helpers

import (
	"fmt"
	"time"
)

func (vld *Validate) ValidateFileInfoIso() error {

	const (
		cycle1        string = "CIC1"
		cycle2        string = "CIC2"
		cycle3        string = "CIC3"
		formatPrev    string = "02/01/2006"
		errorMsgDate  string = `Formato de data inválido: '%s'.`
		errorMsgCycle string = `Ciclo inválido: '%s'. Os ciclos disponíveis são: CIC1, CIC2, CIC3.`
	)

	var err error

	if _, err = time.Parse(formatPrev, vld.FileDate); err != nil {
		return fmt.Errorf(errorMsgDate, vld.FileDate)
	}

	switch vld.FileCycle {
	case cycle1, cycle2, cycle3:
		return nil
	}

	return fmt.Errorf(errorMsgCycle, vld.FileCycle)

}
