package char

import "unicode/utf8"

type Holder struct {
	charList []rune
	charLen  int
}

func (h *Holder) Set(idx int, char rune) {
	h.charList[idx] = char
}

func (h *Holder) ToString() string {
	return string(h.charList)
}

func (h *Holder) ToBytes() []byte {
	size := 0
	for _, r := range h.charList {
		size += utf8.RuneLen(r)
	}

	bs := make([]byte, size)

	count := 0
	for _, r := range h.charList {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs
}

func (h *Holder) GetLen() int {
	return h.charLen
}

func CloneHolder(h *Holder, prefix []rune, suffix []rune) {
	// @todo create new holder
}

func NewHolder(len int, char rune) *Holder {
	chars := make([]rune, len)

	for i := range chars {
		chars[i] = char
	}

	return &Holder{
		charList: chars,
		charLen:  len,
	}
}
