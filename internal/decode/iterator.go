package decode

import (
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/char"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/encode"
)

type Iterator struct {
	encoder  encode.Encoder
	charList []rune
	charLen  int
	minLen   int
	maxLen   int
}

var hash []byte

func (h *Iterator) Run(decHash []byte) (pass string, err error) {
	hash = decHash

	for l := h.minLen; l <= h.maxLen; l++ {
		holder := char.NewHolder(l, h.charList[0])
		pass, err = h.iterateHolder(holder)

		if err == nil {
			return pass, nil
		}
	}

	return "", ErrHashWasNotDecoded
}

func (h Iterator) iterateHolderRune(holder *char.Holder, idx int) (pass string, err error) {
	//charLen := len(h.charList)
	isLastIteration := holder.GetLen() == idx+1

	for charIdx := 0; charIdx < h.charLen; charIdx++ {
		holder.Set(idx, h.charList[charIdx])

		if h.encoder.Match(holder.ToBytes(), hash) {
			return holder.ToString(), nil
		}

		if isLastIteration == false {
			pass, err = h.iterateHolderRune(holder, idx+1)

			if err == nil {
				return pass, nil
			}
		}
	}

	return "", ErrHashWasNotDecoded
}

func (h *Iterator) iterateHolder(holder *char.Holder) (pass string, err error) {
	for runeIdx := 0; runeIdx < holder.GetLen(); runeIdx++ {
		pass, err = h.iterateHolderRune(holder, runeIdx)

		if err == nil {
			return pass, nil
		}
	}

	return "", ErrHashWasNotDecoded
}

func NewIterator(encoder encode.Encoder) *Iterator {
	//chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	//	"abcdefghijklmnopqrstuvwxyz" +
	//	"0123456789")

	chars := []rune("abcdefghijklmnopqrstuvwxyz")

	return &Iterator{
		encoder:  encoder,
		charList: chars,
		charLen:  len(chars),
		minLen:   2,
		maxLen:   5,
	}
}
