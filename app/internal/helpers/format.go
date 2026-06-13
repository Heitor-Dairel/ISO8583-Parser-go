package helpers

import "time"

func RestoreOriginalDate(fileDate string) string {

	const (
		formatPrev string = "02/01/2006"
		formatCurr string = "02012006"
	)

	var date time.Time

	date, _ = time.Parse(formatPrev, fileDate)

	return date.Format(formatCurr)
}
