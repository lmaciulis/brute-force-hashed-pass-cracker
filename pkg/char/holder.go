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

func CloneHolderWithSuffix(h *Holder, suffix []rune) *Holder {
	chars := make([]rune, h.charLen+len(suffix))
	idx := 0

	for _, c := range h.charList {
		chars[idx] = c
		idx++
	}

	for _, c := range suffix {
		chars[idx] = c
		idx++
	}

	return newHolderFromCharList(chars)
}

func CloneHolderWithPrefix(h *Holder, prefix []rune) *Holder {
	chars := make([]rune, h.charLen+len(prefix))
	idx := 0

	for _, c := range prefix {
		chars[idx] = c
		idx++
	}

	for _, c := range h.charList {
		chars[idx] = c
		idx++
	}

	return newHolderFromCharList(chars)
}

func CloneHolder(h *Holder) *Holder {
	chars := make([]rune, h.charLen)

	for i, c := range h.charList {
		chars[i] = c
	}

	return newHolderFromCharList(chars)
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

func newHolderFromCharList(chars []rune) *Holder {
	return &Holder{
		charList: chars,
		charLen:  len(chars),
	}
}
