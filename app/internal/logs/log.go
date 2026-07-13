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

	header, rows = getCsv(log.FileDataParsed)

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

func getCsv(dataParsed []*orderedmap.OrderedMap) (types.HeaderCsv, types.OuterDataCsv) {

	var (
		header types.HeaderCsv    = headerCsv(dataParsed)
		data   types.OuterDataCsv = dataCsv(dataParsed, header)
	)

	return header, data

}

func headerCsv(dataParsed []*orderedmap.OrderedMap) types.HeaderCsv {

	var (
		cache                       map[string]bool = make(map[string]bool, 1500)
		header                      types.HeaderCsv = types.HeaderCsv{"ID"}
		dataElement, subDataElement *orderedmap.OrderedMap
		key                         string
		mti, pds                    any
		isMatch                     bool
	)

	for _, dataElement = range dataParsed {

		mti, _ = dataElement.Get("MTI")
		pds, isMatch = dataElement.Get("PDS")

		if mti == "1240" {
			for _, key = range dataElement.Keys() {
				headerCollect(key, cache, &header)
			}

			if isMatch {
				subDataElement = pds.(*orderedmap.OrderedMap)
				for _, key = range subDataElement.Keys() {
					headerCollect(key, cache, &header)
				}
			}
		}
	}

	headerSort(header)

	return header

}

func headerSort(header types.HeaderCsv) {

	var order map[string]int = map[string]int{"ID": 0, "MTI": 1}

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

}

func headerMatch(key string, hashMap map[string]bool) bool {

	var hashException map[string]bool = map[string]bool{"BMP": true, "BMS": true, "Length": true, "PDS": true}

	return !hashMap[key] && !hashException[key]

}

func headerCollect(key string, hashMap map[string]bool, header *types.HeaderCsv) {

	if headerMatch(key, hashMap) {
		hashMap[key] = true
		*header = append(*header, key)
	}
}

func dataCsv(dataParsed []*orderedmap.OrderedMap, header types.HeaderCsv) types.OuterDataCsv {

	var (
		rows                        types.OuterDataCsv  = make(types.OuterDataCsv, 0, len(dataParsed))
		subRows                     types.NestedDataCsv = make(types.NestedDataCsv, 0, 2000)
		dataElement, subDataElement *orderedmap.OrderedMap
		key                         string
		mti, pds                    any
		id                          int = 1
	)

	for _, dataElement = range dataParsed {

		mti, _ = dataElement.Get("MTI")

		if mti == "1240" {

			subRows = append(subRows, strconv.Itoa(id))

			pds, _ = dataElement.Get("PDS")

			subDataElement = pds.(*orderedmap.OrderedMap)

			for _, key = range header {

				dataCollect(key, dataElement, subDataElement, &subRows)

			}

			rows = append(rows, subRows)

			subRows = make(types.NestedDataCsv, 0, 2000)

			id++

		}

	}

	return rows

}

func dataCollect(key string, dataElement *orderedmap.OrderedMap, subDataElement *orderedmap.OrderedMap, subRows *types.NestedDataCsv) {

	var getOrEmpty func(values any, isExist bool) string
	var value string

	getOrEmpty = func(values any, isExist bool) string {
		if isExist {
			return values.(string)
		}
		return ""
	}

	if !strings.Contains(key, "PDS") && !strings.Contains(key, "ID") {

		value = getOrEmpty(dataElement.Get(key))
		*subRows = append(*subRows, value)

	}

	if strings.Contains(key, "PDS") {

		value = getOrEmpty(subDataElement.Get(key))
		*subRows = append(*subRows, value)

	}
}
