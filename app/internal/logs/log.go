package logs

import (
	"cmp"
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"

	"github.com/bytedance/sonic"
	"github.com/iancoleman/orderedmap"
)

func (log *LogIso) LoggingIso() error {

	var err error

	if err = log.logTxt(); err != nil {
		return err
	}

	if err = log.logCsv(); err != nil {
		return err
	}

	return nil

}

func getDataCsv(dataParsed []*orderedmap.OrderedMap) (types.HeaderCsv, types.OuterDataCsv) {

	var (
		headerCache                 map[string]bool     = make(map[string]bool, 1500)
		headerException             map[string]bool     = map[string]bool{"BMP": true, "BMS": true, "Length": true, "PDS": true}
		order                       map[string]int      = map[string]int{"ID": 0, "MTI": 1}
		header                      types.HeaderCsv     = types.HeaderCsv{"ID"}
		rows                        types.OuterDataCsv  = make(types.OuterDataCsv, 0, len(dataParsed))
		subRows                     types.NestedDataCsv = make(types.NestedDataCsv, 0, 2000)
		dataElement, subDataElement *orderedmap.OrderedMap
		key, value                  string
		mti, pds                    any
		isMatch                     bool
		id                          int = 1
		getOrEmpty                  func(values any, isExist bool) string
	)

	getOrEmpty = func(values any, isExist bool) string {
		if isExist {
			return values.(string)
		}
		return ""
	}

	for _, dataElement = range dataParsed {

		mti, _ = dataElement.Get("MTI")
		pds, isMatch = dataElement.Get("PDS")

		if mti == "1240" {
			for _, key = range dataElement.Keys() {
				if !headerCache[key] && !headerException[key] {
					headerCache[key] = true
					header = append(header, key)
				}
			}

			if isMatch {
				subDataElement = pds.(*orderedmap.OrderedMap)
				for _, key = range subDataElement.Keys() {
					if !headerCache[key] && !headerException[key] {
						headerCache[key] = true
						header = append(header, key)
					}
				}
			}
		}
	}

	slices.SortFunc(header, func(x, y string) int {
		posX, existX := order[x]
		posY, existY := order[y]

		if existX && existY {
			return cmp.Compare(posX, posY)
		}

		if existX {
			return -1
		}

		if existY {
			return 1
		}

		return strings.Compare(x, y)

	})

	for _, dataElement = range dataParsed {

		mti, _ = dataElement.Get("MTI")

		if mti == "1240" {

			subRows = append(subRows, strconv.Itoa(id))

			pds, _ = dataElement.Get("PDS")

			subDataElement = pds.(*orderedmap.OrderedMap)

			for _, key = range header {

				if !strings.Contains(key, "PDS") && !strings.Contains(key, "ID") {

					value = getOrEmpty(dataElement.Get(key))
					subRows = append(subRows, value)

				}

				if strings.Contains(key, "PDS") {

					value = getOrEmpty(subDataElement.Get(key))
					subRows = append(subRows, value)

				}
			}

			rows = append(rows, subRows)

			subRows = make(types.NestedDataCsv, 0, 2000)

			id++

		}

	}

	return header, rows

}

func (log *LogIso) logCsv() error {

	const errorMsg string = "Erro ao criar csv para a data '%s' e ciclo '%s' : %w."
	var (
		header types.HeaderCsv
		rows   types.OuterDataCsv
		file   *os.File
		writer *csv.Writer
		value  []string
		err    error
	)

	header, rows = getDataCsv(log.FileDataParsed)

	if file, err = os.Create(log.PathOutputCsv); err != nil {
		return fmt.Errorf(errorMsg, log.FileDate, log.FileCycle, err)
	}

	defer file.Close()

	writer = csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	writer.Write(header)

	for _, value = range rows {

		writer.Write(value)

	}

	return nil

}
func (log *LogIso) logTxt() error {

	const errorMsg string = "Erro ao criar txt para a data '%s' e ciclo '%s' : %w."
	var (
		bytes types.Data
		err   error
	)

	bytes, _ = sonic.MarshalIndent(log.FileDataParsed, "", "\t")

	if err = os.WriteFile(log.PathOutputTxt, bytes, 0644); err != nil {
		return fmt.Errorf(errorMsg, log.FileDate, log.FileCycle, err)
	}

	return nil
}
