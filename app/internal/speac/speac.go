package speac

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/prefix"
)

var Iso85831993V1 *iso8583.MessageSpec = &iso8583.MessageSpec{
	Name: "Mastercard IPM ISO 8583-1993 EBCDIC",
	Fields: map[int]field.Field{

		0: field.NewString(&field.Spec{
			Length:      4,
			Description: "Message Type Indicator",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		1: field.NewBitmap(&field.Spec{
			Length:      8,
			Description: "Bitmap",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		2: field.NewString(&field.Spec{
			Length:      19,
			Description: "Primary Account number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		3: field.NewString(&field.Spec{
			Length:      6,
			Description: "Processing Code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		4: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount, Transaction",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		5: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount, Reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		6: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount, Cardholder billing",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		7: field.NewString(&field.Spec{
			Length:      10,
			Description: "Date and time, transmission",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		8: field.NewString(&field.Spec{
			Length:      8,
			Description: "Amount, Cardholder billing fee",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		9: field.NewString(&field.Spec{
			Length:      8,
			Description: "Conversion rate, Reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		10: field.NewString(&field.Spec{
			Length:      8,
			Description: "Conversion rate, Cardholder billing",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		11: field.NewString(&field.Spec{
			Length:      6,
			Description: "Systems trace audit number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		12: field.NewString(&field.Spec{
			Length:      12,
			Description: "Date and time, Local transaction",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		13: field.NewString(&field.Spec{
			Length:      4,
			Description: "Date, Effective",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		14: field.NewString(&field.Spec{
			Length:      4,
			Description: "Date, Expiration",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		15: field.NewString(&field.Spec{
			Length:      6,
			Description: "Date, Settlement",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		16: field.NewString(&field.Spec{
			Length:      4,
			Description: "Date, Conversion",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		17: field.NewString(&field.Spec{
			Length:      4,
			Description: "Date, Capture",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		18: field.NewString(&field.Spec{
			Length:      4,
			Description: "Merchant type",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		19: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, Acquiring institution",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		20: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, Primary account number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		21: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, Forwarding institution",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		22: field.NewString(&field.Spec{
			Length:      12,
			Description: "Point of service data code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		23: field.NewString(&field.Spec{
			Length:      3,
			Description: "Card sequence numbe",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		24: field.NewString(&field.Spec{
			Length:      3,
			Description: "Function code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		25: field.NewString(&field.Spec{
			Length:      4,
			Description: "Message reason code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		26: field.NewString(&field.Spec{
			Length:      4,
			Description: "Card acceptor business code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		27: field.NewString(&field.Spec{
			Length:      1,
			Description: "Approval code length",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		28: field.NewString(&field.Spec{
			Length:      6,
			Description: "Date, Reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		29: field.NewString(&field.Spec{
			Length:      3,
			Description: "Reconciliation indicator",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		30: field.NewString(&field.Spec{
			Length:      24,
			Description: "Amounts, original",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		31: field.NewString(&field.Spec{
			Length:      99,
			Description: "Acquirer reference data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		32: field.NewString(&field.Spec{
			Length:      11,
			Description: "Acquirer institution identification code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		33: field.NewString(&field.Spec{
			Length:      11,
			Description: "Forwarding institution identification code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		34: field.NewString(&field.Spec{
			Length:      28,
			Description: "Primary account number, extended",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		35: field.NewString(&field.Spec{
			Length:      37,
			Description: "Track 2 data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		36: field.NewString(&field.Spec{
			Length:      104,
			Description: "Track 3 data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		37: field.NewString(&field.Spec{
			Length:      12,
			Description: "Retrieval reference number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		38: field.NewString(&field.Spec{
			Length:      6,
			Description: "Approval code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		39: field.NewString(&field.Spec{
			Length:      3,
			Description: "Action code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		40: field.NewString(&field.Spec{
			Length:      3,
			Description: "Service code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		41: field.NewString(&field.Spec{
			Length:      8,
			Description: "Card acceptor terminal identification",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		42: field.NewString(&field.Spec{
			Length:      15,
			Description: "Card acceptor identification code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		43: field.NewString(&field.Spec{
			Length:      99,
			Description: "Card acceptor name/location",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		44: field.NewString(&field.Spec{
			Length:      99,
			Description: "Additional response data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		45: field.NewString(&field.Spec{
			Length:      76,
			Description: "Track 1 data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		46: field.NewString(&field.Spec{
			Length:      204,
			Description: "Amounts, Fees",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		47: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional data - national",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		48: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional data - private",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		49: field.NewString(&field.Spec{
			Length:      3,
			Description: "Currency code, Transaction",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		50: field.NewString(&field.Spec{
			Length:      3,
			Description: "Currency code, Reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		51: field.NewString(&field.Spec{
			Length:      3,
			Description: "Currency code, Cardholder billing",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		52: field.NewString(&field.Spec{
			Length:      8,
			Description: "Personal identification number [PIN] data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		53: field.NewString(&field.Spec{
			Length:      48,
			Description: "Security related control information",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		54: field.NewString(&field.Spec{
			Length:      120,
			Description: "Amounts, additional",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		55: field.NewBinary(&field.Spec{
			Length:      255,
			Description: "IC card system related data",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.LLL,
		}),

		56: field.NewString(&field.Spec{
			Length:      35,
			Description: "Original data elements",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		57: field.NewString(&field.Spec{
			Length:      3,
			Description: "Authorization life cycle cod",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		58: field.NewString(&field.Spec{
			Length:      11,
			Description: "Authorizing agent institution Id Code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		59: field.NewString(&field.Spec{
			Length:      999,
			Description: "Transport data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		60: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		61: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		62: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		63: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		64: field.NewBinary(&field.Spec{
			Length:      8,
			Description: "Message authentication code field",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		65: field.NewBinary(&field.Spec{
			Length:      8,
			Description: "Reserved for ISO use",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		66: field.NewString(&field.Spec{
			Length:      204,
			Description: "Amounts, original fees",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		67: field.NewString(&field.Spec{
			Length:      2,
			Description: "Extended payment data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		68: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, receiving institution",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		69: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, settlement institution",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		70: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, authorizing agent Inst.",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		71: field.NewString(&field.Spec{
			Length:      8,
			Description: "Message number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		72: field.NewString(&field.Spec{
			Length:      999,
			Description: "Data record",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		73: field.NewString(&field.Spec{
			Length:      6,
			Description: "Date, action",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		74: field.NewString(&field.Spec{
			Length:      10,
			Description: "Credits, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		75: field.NewString(&field.Spec{
			Length:      10,
			Description: "Credits, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		76: field.NewString(&field.Spec{
			Length:      10,
			Description: "Debits, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		77: field.NewString(&field.Spec{
			Length:      10,
			Description: "Debits, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		78: field.NewString(&field.Spec{
			Length:      10,
			Description: "Transfer, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		79: field.NewString(&field.Spec{
			Length:      10,
			Description: "Transfer, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		80: field.NewString(&field.Spec{
			Length:      10,
			Description: "Inquiries, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		81: field.NewString(&field.Spec{
			Length:      10,
			Description: "Authorizations, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		82: field.NewString(&field.Spec{
			Length:      10,
			Description: "Inquiries, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		83: field.NewString(&field.Spec{
			Length:      10,
			Description: "Payments, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		84: field.NewString(&field.Spec{
			Length:      10,
			Description: "Payments, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		85: field.NewString(&field.Spec{
			Length:      10,
			Description: "Fee collections, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		86: field.NewString(&field.Spec{
			Length:      16,
			Description: "Credits, amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		87: field.NewString(&field.Spec{
			Length:      16,
			Description: "Credits, reversal amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		88: field.NewString(&field.Spec{
			Length:      16,
			Description: "Debits, amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		89: field.NewString(&field.Spec{
			Length:      16,
			Description: "Debits, reversal amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		90: field.NewString(&field.Spec{
			Length:      10,
			Description: "Authorizations, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		91: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, transaction Dest. Inst.",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		92: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, transaction Orig. Inst.",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		93: field.NewString(&field.Spec{
			Length:      11,
			Description: "Transaction Dest. Inst. Id code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		94: field.NewString(&field.Spec{
			Length:      11,
			Description: "Transaction Orig. Inst. Id code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		95: field.NewString(&field.Spec{
			Length:      99,
			Description: "Card issuer reference data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		96: field.NewString(&field.Spec{
			Length:      999,
			Description: "Key management data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		97: field.NewString(&field.Spec{
			Length:      17,
			Description: "Amount, Net reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		98: field.NewString(&field.Spec{
			Length:      25,
			Description: "Payee",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		99: field.NewString(&field.Spec{
			Length:      11,
			Description: "Settlement institution Id code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		100: field.NewString(&field.Spec{
			Length:      11,
			Description: "Receiving institution Id code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		101: field.NewString(&field.Spec{
			Length:      17,
			Description: "File name",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		102: field.NewString(&field.Spec{
			Length:      28,
			Description: "Account identification 1",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		103: field.NewString(&field.Spec{
			Length:      28,
			Description: "Account identification 2",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		104: field.NewString(&field.Spec{
			Length:      100,
			Description: "Transaction description",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		105: field.NewString(&field.Spec{
			Length:      16,
			Description: "Credits, Chargeback amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		106: field.NewString(&field.Spec{
			Length:      16,
			Description: "Debits, Chargeback amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		107: field.NewString(&field.Spec{
			Length:      10,
			Description: "Credits, Chargeback number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		108: field.NewString(&field.Spec{
			Length:      10,
			Description: "Debits, Chargeback number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		109: field.NewString(&field.Spec{
			Length:      84,
			Description: "Credits, Fee amounts",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		110: field.NewString(&field.Spec{
			Length:      84,
			Description: "Debits, Fee amounts",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		111: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		112: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		113: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		114: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		115: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		116: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		117: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		118: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		119: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		120: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		121: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		122: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		123: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		124: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		125: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		126: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		127: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		128: field.NewBinary(&field.Spec{
			Length:      8,
			Description: "Message authentication code field",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.Fixed,
		}),
	},
}

var Iso85831993V2 *iso8583.MessageSpec = &iso8583.MessageSpec{
	Name: "Mastercard IPM ISO 8583-1993 EBCDIC",
	Fields: map[int]field.Field{

		0: field.NewString(&field.Spec{
			Length:      4,
			Description: "Message Type Indicator",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		1: field.NewBitmap(&field.Spec{
			Length:      8,
			Description: "Bitmap",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		2: field.NewString(&field.Spec{
			Length:      19,
			Description: "Primary Account number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		3: field.NewString(&field.Spec{
			Length:      6,
			Description: "Processing Code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		4: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount, Transaction",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		5: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount, Reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		6: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount, Cardholder billing",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		7: field.NewString(&field.Spec{
			Length:      10,
			Description: "Date and time, transmission",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		8: field.NewString(&field.Spec{
			Length:      8,
			Description: "Amount, Cardholder billing fee",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		9: field.NewString(&field.Spec{
			Length:      8,
			Description: "Conversion rate, Reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		10: field.NewString(&field.Spec{
			Length:      8,
			Description: "Conversion rate, Cardholder billing",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		11: field.NewString(&field.Spec{
			Length:      6,
			Description: "Systems trace audit number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		12: field.NewString(&field.Spec{
			Length:      12,
			Description: "Date and time, Local transaction",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		13: field.NewString(&field.Spec{
			Length:      4,
			Description: "Date, Effective",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		14: field.NewString(&field.Spec{
			Length:      4,
			Description: "Date, Expiration",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		15: field.NewString(&field.Spec{
			Length:      4,
			Description: "Date, Settlement",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		16: field.NewString(&field.Spec{
			Length:      4,
			Description: "Date, Conversion",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		17: field.NewString(&field.Spec{
			Length:      4,
			Description: "Date, Capture",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		18: field.NewString(&field.Spec{
			Length:      4,
			Description: "Merchant type",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		19: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, Acquiring institution",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		20: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, Primary account number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		21: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, Forwarding institution",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		22: field.NewString(&field.Spec{
			Length:      12,
			Description: "Point of service data code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		23: field.NewString(&field.Spec{
			Length:      3,
			Description: "Card sequence numbe",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		24: field.NewString(&field.Spec{
			Length:      3,
			Description: "Function code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		25: field.NewString(&field.Spec{
			Length:      4,
			Description: "Message reason code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		26: field.NewString(&field.Spec{
			Length:      4,
			Description: "Card acceptor business code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		27: field.NewString(&field.Spec{
			Length:      1,
			Description: "Approval code length",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		28: field.NewString(&field.Spec{
			Length:      9,
			Description: "Date, Reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		29: field.NewString(&field.Spec{
			Length:      9,
			Description: "Reconciliation indicator",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		30: field.NewString(&field.Spec{
			Length:      24,
			Description: "Amounts, original",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		31: field.NewString(&field.Spec{
			Length:      23,
			Description: "Acquirer reference data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		32: field.NewString(&field.Spec{
			Length:      11,
			Description: "Acquirer institution identification code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		33: field.NewString(&field.Spec{
			Length:      11,
			Description: "Forwarding institution identification code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		34: field.NewString(&field.Spec{
			Length:      28,
			Description: "Primary account number, extended",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		35: field.NewString(&field.Spec{
			Length:      37,
			Description: "Track 2 data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		36: field.NewString(&field.Spec{
			Length:      104,
			Description: "Track 3 data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		37: field.NewString(&field.Spec{
			Length:      12,
			Description: "Retrieval reference number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		38: field.NewString(&field.Spec{
			Length:      6,
			Description: "Approval code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		39: field.NewString(&field.Spec{
			Length:      2,
			Description: "Action code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		40: field.NewString(&field.Spec{
			Length:      3,
			Description: "Service code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		41: field.NewString(&field.Spec{
			Length:      8,
			Description: "Card acceptor terminal identification",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		42: field.NewString(&field.Spec{
			Length:      15,
			Description: "Card acceptor identification code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		43: field.NewString(&field.Spec{
			Length:      99,
			Description: "Card acceptor name/location",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		44: field.NewString(&field.Spec{
			Length:      25,
			Description: "Additional response data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		45: field.NewString(&field.Spec{
			Length:      76,
			Description: "Track 1 data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		46: field.NewString(&field.Spec{
			Length:      999,
			Description: "Amounts, Fees",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		47: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional data - national",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		48: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional data - private",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		49: field.NewString(&field.Spec{
			Length:      3,
			Description: "Currency code, Transaction",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		50: field.NewString(&field.Spec{
			Length:      3,
			Description: "Currency code, Reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		51: field.NewString(&field.Spec{
			Length:      3,
			Description: "Currency code, Cardholder billing",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		52: field.NewString(&field.Spec{
			Length:      8,
			Description: "Personal identification number [PIN] data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		53: field.NewString(&field.Spec{
			Length:      16,
			Description: "Security related control information",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		54: field.NewString(&field.Spec{
			Length:      240,
			Description: "Amounts, additional",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		55: field.NewBinary(&field.Spec{
			Length:      999,
			Description: "IC card system related data",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.LLL,
		}),

		56: field.NewString(&field.Spec{
			Length:      999,
			Description: "Original data elements",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		57: field.NewString(&field.Spec{
			Length:      999,
			Description: "Authorization life cycle cod",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		58: field.NewString(&field.Spec{
			Length:      999,
			Description: "Authorizing agent institution Id Code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		59: field.NewString(&field.Spec{
			Length:      999,
			Description: "Transport data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		60: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		61: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		62: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		63: field.NewString(&field.Spec{
			Length:      16,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		64: field.NewBinary(&field.Spec{
			Length:      8,
			Description: "Message authentication code field",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		65: field.NewBinary(&field.Spec{
			Length:      8,
			Description: "Reserved for ISO use",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		66: field.NewString(&field.Spec{
			Length:      1,
			Description: "Amounts, original fees",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		67: field.NewString(&field.Spec{
			Length:      2,
			Description: "Extended payment data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		68: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, receiving institution",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		69: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, settlement institution",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		70: field.NewString(&field.Spec{
			Length:      3,
			Description: "Country code, authorizing agent Inst.",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		71: field.NewString(&field.Spec{
			Length:      8,
			Description: "Message number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		72: field.NewString(&field.Spec{
			Length:      999,
			Description: "Data record",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		73: field.NewString(&field.Spec{
			Length:      6,
			Description: "Date, action",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		74: field.NewString(&field.Spec{
			Length:      10,
			Description: "Credits, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		75: field.NewString(&field.Spec{
			Length:      10,
			Description: "Credits, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		76: field.NewString(&field.Spec{
			Length:      10,
			Description: "Debits, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		77: field.NewString(&field.Spec{
			Length:      10,
			Description: "Debits, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		78: field.NewString(&field.Spec{
			Length:      10,
			Description: "Transfer, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		79: field.NewString(&field.Spec{
			Length:      10,
			Description: "Transfer, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		80: field.NewString(&field.Spec{
			Length:      10,
			Description: "Inquiries, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		81: field.NewString(&field.Spec{
			Length:      10,
			Description: "Authorizations, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		82: field.NewString(&field.Spec{
			Length:      12,
			Description: "Inquiries, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		83: field.NewString(&field.Spec{
			Length:      12,
			Description: "Payments, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		84: field.NewString(&field.Spec{
			Length:      12,
			Description: "Payments, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		85: field.NewString(&field.Spec{
			Length:      12,
			Description: "Fee collections, number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		86: field.NewString(&field.Spec{
			Length:      16,
			Description: "Credits, amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		87: field.NewString(&field.Spec{
			Length:      16,
			Description: "Credits, reversal amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		88: field.NewString(&field.Spec{
			Length:      16,
			Description: "Debits, amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		89: field.NewString(&field.Spec{
			Length:      16,
			Description: "Debits, reversal amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		90: field.NewString(&field.Spec{
			Length:      42,
			Description: "Authorizations, reversal number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		91: field.NewString(&field.Spec{
			Length:      1,
			Description: "Country code, transaction Dest. Inst.",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		92: field.NewString(&field.Spec{
			Length:      2,
			Description: "Country code, transaction Orig. Inst.",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		93: field.NewString(&field.Spec{
			Length:      11,
			Description: "Transaction Dest. Inst. Id code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		94: field.NewString(&field.Spec{
			Length:      11,
			Description: "Transaction Orig. Inst. Id code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		95: field.NewString(&field.Spec{
			Length:      10,
			Description: "Card issuer reference data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		96: field.NewString(&field.Spec{
			Length:      8,
			Description: "Key management data",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		97: field.NewString(&field.Spec{
			Length:      17,
			Description: "Amount, Net reconciliation",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		98: field.NewString(&field.Spec{
			Length:      25,
			Description: "Payee",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),

		99: field.NewString(&field.Spec{
			Length:      11,
			Description: "Settlement institution Id code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		100: field.NewString(&field.Spec{
			Length:      11,
			Description: "Receiving institution Id code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		101: field.NewString(&field.Spec{
			Length:      17,
			Description: "File name",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		102: field.NewString(&field.Spec{
			Length:      28,
			Description: "Account identification 1",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		103: field.NewString(&field.Spec{
			Length:      28,
			Description: "Account identification 2",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LL,
		}),

		104: field.NewString(&field.Spec{
			Length:      100,
			Description: "Transaction description",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		105: field.NewString(&field.Spec{
			Length:      999,
			Description: "Credits, Chargeback amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		106: field.NewString(&field.Spec{
			Length:      999,
			Description: "Debits, Chargeback amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		107: field.NewString(&field.Spec{
			Length:      999,
			Description: "Credits, Chargeback number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		108: field.NewString(&field.Spec{
			Length:      999,
			Description: "Debits, Chargeback number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		109: field.NewString(&field.Spec{
			Length:      999,
			Description: "Credits, Fee amounts",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		110: field.NewString(&field.Spec{
			Length:      999,
			Description: "Debits, Fee amounts",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		111: field.NewString(&field.Spec{
			Length:      12,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		112: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		113: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		114: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		115: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for ISO use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		116: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		117: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		118: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		119: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		120: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		121: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		122: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for national use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		123: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		124: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		125: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		126: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		127: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved for private use",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.LLL,
		}),

		128: field.NewBinary(&field.Spec{
			Length:      8,
			Description: "Message authentication code field",
			Enc:         encoding.Binary,
			Pref:        prefix.EBCDIC.Fixed,
		}),
	},
}
