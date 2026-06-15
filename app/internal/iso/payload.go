package iso

import (
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/parse"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/speac"
	"github.com/Heitor-Dairel/ISO8583-Parser-go/app/internal/types"

	"github.com/iancoleman/orderedmap"
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/field"
)

func extractIsoPayload(raw []byte, index, lenRaw int) (types.Data, int) {
	var (
		indexStart, indexCurr     int        = index, index
		payload                   types.Data = make(types.Data, 0, 5000)
		segId, segLen, payloadLen int
	)

	for {

		if indexCurr+4 > lenRaw {
			break
		}

		segId = int(raw[indexCurr+2] & 0xFF)
		segLen = int(raw[indexCurr]&0xFF)<<8 | int(raw[indexCurr+1]&0xFF)

		payloadLen = segLen - 4

		payload = append(payload, raw[indexCurr+4:indexCurr+4+payloadLen]...)

		indexCurr += 4 + payloadLen

		if segId == 0 {
			break
		}

		if indexCurr+2 < lenRaw && raw[indexCurr+2] == 0 {
			break
		}

	}

	return payload, indexCurr - indexStart

}

func (iso *ISO8583) fileIsoPayload() error {
	var (
		index, consumed int
		lenRaw          int        = len(iso.FilePathIso.FileData)
		payload         types.Data = make(types.Data, 0, 5000)
		msg             *iso8583.Message
		field           map[int]field.Field = make(map[int]field.Field, 129)
		parseOrdMap     *orderedmap.OrderedMap
		err             error
	)

	for index < lenRaw {

		payload, consumed = extractIsoPayload(iso.FilePathIso.FileData, index, lenRaw)

		index += consumed

		msg = iso8583.NewMessage(speac.Iso85831993V1)

		if err = msg.Unpack(payload); err != nil {
			return err
		}

		field = msg.GetFields()

		parseOrdMap = parse.ParseBeautify(field, consumed)

		iso.FileDataParsed = append(iso.FileDataParsed, parseOrdMap)

	}

	iso.ParsedLines = len(iso.FileDataParsed)

	return nil

}
