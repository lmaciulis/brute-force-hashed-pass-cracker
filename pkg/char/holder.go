package char

import "unicode/utf8"

type Holder struct {
	charList []rune
	charLen  int
}

// Set replace holder's char with new one, by providing index and new char
func (h *Holder) Set(idx int, char rune) {
	h.charList[idx] = char
}

// ToString concatenates holder's runes to string
func (h *Holder) ToString() string {
	return string(h.charList)
}

// ToBytes returns holder's runes in bytes representation as array of bytes
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

// GetLen returns holder's length
func (h *Holder) GetLen() int {
	return h.charLen
}

// CloneHolderWithSuffix clone provided holder, but with appended suffix
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

// CloneHolderWithPrefix clone provided holder, but with prepended prefix
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

// NewHolder creates new holder identical to provided one
func CloneHolder(h *Holder) *Holder {
	chars := make([]rune, h.charLen)

	for i, c := range h.charList {
		chars[i] = c
	}

	return newHolderFromCharList(chars)
}

// NewHolder crates new chars holder with given length and fills every holder's char with same provided character
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

// newHolderFromCharList creates new chars holder, by providing complete chars list
func newHolderFromCharList(chars []rune) *Holder {
	return &Holder{
		charList: chars,
		charLen:  len(chars),
	}
}
