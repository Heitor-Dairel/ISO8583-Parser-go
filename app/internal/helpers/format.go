package helpers

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RestoreOriginalDate(fileDate string) string {

	const (
		formatPrev string = "02/01/2006"
		formatCurr string = "02012006"
	)

	var date time.Time

	date, _ = time.Parse(formatPrev, fileDate)

	return date.Format(formatCurr)
}

func FormatDateTime(fileName string) string {

	const (
		formatPrev string = "02012006150405"
		formatCurr string = "02/01/2006 15:04:05"
	)

	var (
		date     time.Time
		parts    []string = strings.Split(fileName, "_")
		lenParts int      = len(parts)
	)

	date, _ = time.Parse(formatPrev, parts[lenParts-2]+parts[lenParts-1])
	return date.Format(formatCurr)
}

func FormatSize(filSize int64) string {

	const (
		mb int64 = 1024 * 1024
		kb int64 = 1024
	)

	var fileSizeFloat float64 = float64(filSize)

	if filSize >= mb {
		return fmt.Sprintf("%.2f MB", fileSizeFloat/float64(mb))
	}

	if filSize >= kb {
		return fmt.Sprintf("%.2f KB", fileSizeFloat/float64(kb))
	}

	return fmt.Sprintf("%d Bytes", filSize)

}

func FormatLines(fileLines int) string {
	var p *message.Printer = message.NewPrinter(language.BrazilianPortuguese)
	return p.Sprintf("%d", fileLines)
}
