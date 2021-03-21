package char

import "unicode/utf8"

type Holder struct {
	chars []rune
	count int
}

func (h *Holder) Set(idx int, char rune) {
	h.chars[idx] = char
}

func (h *Holder) ToString() string {
	return string(h.chars)
}

func (h *Holder) ToBytes() []byte {
	size := 0
	for _, r := range h.chars {
		size += utf8.RuneLen(r)
	}

	bs := make([]byte, size)

	count := 0
	for _, r := range h.chars {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs
}

func (h *Holder) GetLen() int {
	return h.count
}

func NewHolder(len int, char rune) *Holder {
	chars := make([]rune, len)

	for i := range chars {
		chars[i] = char
	}

	return &Holder{
		chars: chars,
		count: len,
	}
}
