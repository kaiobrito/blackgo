package deck

import "strconv"

type CardNumber int64

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
	C11 //J
	C12 //Q
	C13 //K
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
		C11, //J
		C12, //Q
		C13, //K
	}
}

func (cardNumber CardNumber) Value() string {
	switch cardNumber {
	case CA:
		return "A"
	case C11:
		return "J"
	case C12:
		return "Q"
	case C13:
		return "K"
	}
	return strconv.Itoa(int(cardNumber))
}
