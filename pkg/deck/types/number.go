package deck

import (
	"encoding/json"
	"strconv"
)

type CardNumber int64

func (a CardNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Value())
}

const (
	CA CardNumber = iota
	C1
	C2
	C3
	C4
	C5
	C6
	C7
	C8
	C9
	C10
	CJ //J
	CQ //Q
	CK //K
)

func AllCardNumbers() []CardNumber {
	return []CardNumber{
		CA,
		C1,
		C2,
		C3,
		C4,
		C5,
		C6,
		C7,
		C8,
		C9,
		C10,
		CJ, //J
		CQ, //Q
		CK, //K
	}
}

func (cardNumber CardNumber) Value() string {
	switch cardNumber {
	case CA:
		return "A"
	case CJ:
		return "J"
	case CQ:
		return "Q"
	case CK:
		return "K"
	}
	return strconv.Itoa(int(cardNumber))
}
