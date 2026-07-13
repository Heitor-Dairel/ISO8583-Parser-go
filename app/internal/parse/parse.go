package parse

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/logs"
	"github.com/iancoleman/orderedmap"
	"github.com/moov-io/iso8583/field"
)

func parseDE048(data string) *orderedmap.OrderedMap {

	const (
		lenTagField, lenPdsField int    = 4, 7
		errorMsg                 string = "Erro ao criar os PDS do data element DE048."
	)

	var (
		lenData                      int = len(data)
		tag, value                   string
		lenPds, sepParse, start, end int
		err                          error
		pds048                       *orderedmap.OrderedMap = orderedmap.New()
	)

	for sepParse < lenData {
		tag = fmt.Sprintf("PDS%s", data[sepParse:sepParse+lenTagField])
		lenPds, err = strconv.Atoi(data[sepParse+lenTagField : sepParse+lenPdsField])
		if err != nil {
			logs.Loggers.Fatal(errorMsg)
		}
		start = sepParse + lenPdsField
		end = sepParse + lenPdsField + lenPds
		value = data[start:end]
		pds048.Set(tag, value)
		sepParse = end
	}
	return pds048

}

func orderKeys(data map[int]field.Field) ordKeys {

	var (
		fieldBitPrimary   []int          = make([]int, 0, 65)
		fieldBitSecondary []int          = make([]int, 0, 65)
		keys              []int          = make([]int, 0, len(data))
		dataElement       map[int]string = make(map[int]string, 129)
		id                int
	)

	for id = range data {
		keys = append(keys, id)
	}

	slices.Sort(keys)

	for _, id = range keys {

		dataElement[id] = fmt.Sprintf("DE%03d", id)

		if id >= 1 && id <= 64 {
			fieldBitPrimary = append(fieldBitPrimary, id)
		}

		if id >= 65 {
			fieldBitSecondary = append(fieldBitSecondary, id)
		}

	}

	return ordKeys{keys: keys, fieldBitPrimary: fieldBitPrimary, fieldBitSecondary: fieldBitSecondary, dataElement: dataElement}
}

func ParseBeautify(data map[int]field.Field, lenData int) *orderedmap.OrderedMap {

	var (
		ordKeys       ordKeys
		parseOrdMap   *orderedmap.OrderedMap = orderedmap.New()
		de048, values string
		parseField    field.Field
		id            int
		err           error
		parseSet      func(key string, value any) = parseOrdMap.Set
	)

	ordKeys = orderKeys(data)

	for _, id = range ordKeys.keys {

		parseField, _ = data[id]

		values, err = parseField.String()

		if err != nil {
			continue
		}

		if id == 0 {

			parseSet("MTI", values)
			parseSet("BMP", ordKeys.fieldBitPrimary)
			parseSet("BMS", ordKeys.fieldBitSecondary)

		}

		if id == 48 {

			de048 = values

		}

		if id > 1 {

			parseSet(ordKeys.dataElement[id], values)

		}

	}

	if de048 != "" {

		parseSet("PDS", parseDE048(de048))

	}

	parseSet("Length", lenData)

	return parseOrdMap

}
