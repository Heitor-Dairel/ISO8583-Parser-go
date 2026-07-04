package iso

import (
	"fmt"

	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/helpers"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/parse"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/speac"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"

	"github.com/iancoleman/orderedmap"
	"github.com/moov-io/iso8583"
)

func extractIsoPayload(data []byte, index, lenData int) (types.Data, int) {
	var (
		indexStart, indexCurr     int        = index, index
		payload                   types.Data = make(types.Data, 0, 5000)
		segId, segLen, payloadLen int
	)

	for {

		if indexCurr+4 > lenData {
			break
		}

		segId = int(data[indexCurr+2] & 0xFF)
		segLen = int(data[indexCurr]&0xFF)<<8 | int(data[indexCurr+1]&0xFF)

		payloadLen = segLen - 4

		payload = append(payload, data[indexCurr+4:indexCurr+4+payloadLen]...)

		indexCurr += 4 + payloadLen

		if segId == 0 {
			break
		}

		if indexCurr+2 < lenData && data[indexCurr+2] == 0 {
			break
		}

	}

	return payload, indexCurr - indexStart

}

func (iso *ISO8583) fileIsoPayload() error {

	const errorMsgUnpackIso string = "Erro ao realizar parse na linha '%d' : %w."

	var (
		index, consumed, msgCount int
		lenData                   int = len(iso.FilePathIso.FileData)
		payload                   types.Data
		msg                       *iso8583.Message
		parseOrdMap               *orderedmap.OrderedMap
		err                       error
	)

	msg = iso8583.NewMessage(speac.Iso85831993V1)

	for index < lenData {

		payload, consumed = extractIsoPayload(iso.FilePathIso.FileData, index, lenData)

		index += consumed

		if err = msg.Unpack(payload); err != nil {
			return fmt.Errorf(errorMsgUnpackIso, msgCount, err)
		}

		parseOrdMap = parse.ParseBeautify(msg.GetFields(), consumed)

		iso.FileDataParsed = append(iso.FileDataParsed, parseOrdMap)

		msgCount++

	}

	iso.ParsedLines = helpers.FormatLines(msgCount)

	return nil

}
