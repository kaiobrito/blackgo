package deck

import (
	"strconv"
)

type CardNumber string

const (
	CA  CardNumber = "A"
	C1  CardNumber = "1"
	C2  CardNumber = "2"
	C3  CardNumber = "3"
	C4  CardNumber = "4"
	C5  CardNumber = "5"
	C6  CardNumber = "6"
	C7  CardNumber = "7"
	C8  CardNumber = "8"
	C9  CardNumber = "9"
	C10 CardNumber = "10"
	CJ  CardNumber = "J"
	CQ  CardNumber = "Q"
	CK  CardNumber = "K"
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
func (c CardNumber) Value() int {
	switch c {
	case CA:
		return 0
	case CJ:
		return 11
	case CQ:
		return 12
	case CK:
		return 13
	}

	if value, err := strconv.Atoi(string(c)); err == nil {
		return value
	}
	return -1
}
