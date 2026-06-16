package parse

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/iancoleman/orderedmap"
	"github.com/moov-io/iso8583/field"
)

func parseDE048(raw string) *orderedmap.OrderedMap {

	const (
		lenTagField, lenPdsField int    = 4, 7
		errorMsg                 string = `Erro ao criar os PDS do data element DE048.`
	)

	var (
		lenRaw                       int = len(raw)
		tag, value                   string
		lenPds, sepParse, start, end int
		err                          error
		pds048                       *orderedmap.OrderedMap = orderedmap.New()
	)

	for sepParse < lenRaw {
		tag = fmt.Sprintf("PDS%s", raw[sepParse:sepParse+lenTagField])
		lenPds, err = strconv.Atoi(raw[sepParse+lenTagField : sepParse+lenPdsField])
		if err != nil {
			log.Fatalln(errorMsg)
		}
		start = sepParse + lenPdsField
		end = sepParse + lenPdsField + lenPds
		value = raw[start:end]
		pds048.Set(tag, value)
		sepParse = end
	}
	return pds048

}

func ParseBeautify(raw map[int]field.Field, lenRaw int) *orderedmap.OrderedMap {

	var (
		fieldBitPrimary   []int                  = make([]int, 0, 65)
		fieldBitSecundary []int                  = make([]int, 0, 65)
		keys              []int                  = make([]int, 0, len(raw))
		parseOrdMap       *orderedmap.OrderedMap = orderedmap.New()
		dataElement       map[int]string         = make(map[int]string, 129)
		de048, values     string
		parseField        field.Field
		id                int
		err               error
		parseSet          func(key string, value any) = parseOrdMap.Set
	)

	for id = range raw {
		keys = append(keys, id)
	}

	slices.Sort(keys)

	for _, id = range keys {

		dataElement[id] = fmt.Sprintf("DE%03d", id)

		if id >= 1 && id <= 64 {
			fieldBitPrimary = append(fieldBitPrimary, id)
		}

		if id >= 65 {
			fieldBitSecundary = append(fieldBitSecundary, id)
		}

	}

	for _, id = range keys {

		parseField, _ = raw[id]

		values, err = parseField.String()

		if err != nil {
			continue
		}

		if id == 0 {

			parseSet("MTI", values)
			parseSet("BMP", fieldBitPrimary)
			parseSet("BMS", fieldBitSecundary)

		}

		if id == 48 {

			de048 = values

		}

		if id > 1 {

			parseSet(dataElement[id], values)

		}

	}

	if de048 != "" {

		parseSet("PDS", parseDE048(de048))

	}

	parseSet("Length", lenRaw)

	return parseOrdMap

}
