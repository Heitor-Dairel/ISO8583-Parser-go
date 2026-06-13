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

func (log *Log) LoggingIso() error {

	var err error

	err = log.logTxt()

	if err != nil {
		return err
	}

	err = log.logCsv()

	if err != nil {
		return err
	}

	return nil

}

func getDataCsv(keys []*orderedmap.OrderedMap) (types.HeaderCsv, types.OuterDataCsv) {

	var (
		headerCache     map[string]bool     = make(map[string]bool)
		headerException map[string]bool     = map[string]bool{"BMP": true, "BMS": true, "Length": true, "PDS": true}
		order           map[string]int      = map[string]int{"ID": 0, "MTI": 1}
		header          types.HeaderCsv     = types.HeaderCsv{"ID"}
		dataCsv         types.OuterDataCsv  = make(types.OuterDataCsv, 0, 7000)
		subDataCsv      types.NestedDataCsv = make(types.NestedDataCsv, 0, 1000)
		row, subRow     *orderedmap.OrderedMap
		key, valueCurr  string
		mti, pds        any
		isMatch         bool
		id              int = 1
		getOrEmpty      func(values any, isExist bool) string
	)

	getOrEmpty = func(values any, isExist bool) string {
		if isExist {
			return values.(string)
		}
		return ""
	}

	for i := 0; i < len(keys); i++ {

		row = keys[i]
		mti, _ = row.Get("MTI")
		pds, isMatch = row.Get("PDS")

		if mti == "1240" {
			for _, key = range row.Keys() {
				if !headerCache[key] && !headerException[key] {
					headerCache[key] = true
					header = append(header, key)
				}
			}

			if isMatch {
				subRow = pds.(*orderedmap.OrderedMap)
				for _, key = range subRow.Keys() {
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

	for i := 0; i < len(keys); i++ {

		row = keys[i]

		mti, _ = row.Get("MTI")

		if mti == "1240" {

			subDataCsv = append(subDataCsv, strconv.Itoa(id))

			pds, _ = row.Get("PDS")

			subRow = pds.(*orderedmap.OrderedMap)

			for _, key = range header {

				if !strings.Contains(key, "PDS") && !strings.Contains(key, "ID") {

					valueCurr = getOrEmpty(row.Get(key))
					subDataCsv = append(subDataCsv, valueCurr)

				}

				if strings.Contains(key, "PDS") {

					valueCurr = getOrEmpty(subRow.Get(key))
					subDataCsv = append(subDataCsv, valueCurr)

				}
			}

			dataCsv = append(dataCsv, subDataCsv)

			subDataCsv = types.NestedDataCsv{}

			id++

		}

	}

	return header, dataCsv

}

func (lg *Log) logCsv() error {

	const errorMsg string = `Erro ao criar arquivo csv: %w`
	var (
		header  types.HeaderCsv
		dataCsv types.OuterDataCsv
		file    *os.File
		writer  *csv.Writer
		val     []string
		err     error
	)

	header, dataCsv = getDataCsv(lg.FileDataParse)

	file, err = os.Create(lg.PathCsv)

	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}

	defer file.Close()

	writer = csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	writer.Write(header)

	for _, val = range dataCsv {

		writer.Write(val)

	}

	return nil

}
func (lg *Log) logTxt() error {

	const errorMsg string = `Erro ao criar arquivo txt: %w`
	var (
		bytes types.Data
		err   error
	)

	bytes, _ = sonic.MarshalIndent(lg.FileDataParse, "", "\t")

	err = os.WriteFile(lg.PathTxt, bytes, 0644)

	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}

	return nil
}
