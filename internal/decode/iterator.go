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

func (i *Iterator) Run(decHash []byte) (pass string, err error) {
	hash = decHash

	for hLen := i.minLen; hLen <= i.maxLen; hLen++ {
		holder := char.NewHolder(hLen, i.getFirstChar())
		pass, err = i.iterateHolder(holder)

		if err == nil {
			return pass, nil
		}
	}

	return "", ErrHashWasNotDecoded
}

func (i *Iterator) iterateHolder(holder *char.Holder) (pass string, err error) {
	for hIdx := 0; hIdx < holder.GetLen(); hIdx++ {
		pass, err = i.iterateHoldersRune(holder, hIdx)

		if err == nil {
			return pass, nil
		}
	}

	return "", ErrHashWasNotDecoded
}

func (i Iterator) iterateHoldersRune(holder *char.Holder, idx int) (pass string, err error) {
	isLastIteration := holder.GetLen() == idx+1

	for charIdx := 0; charIdx < i.charLen; charIdx++ {
		holder.Set(idx, i.getChar(charIdx))

		if i.encoder.Match(holder.ToBytes(), hash) {
			return holder.ToString(), nil
		}

		if isLastIteration == false {
			pass, err = i.iterateHoldersRune(holder, idx+1)

			if err == nil {
				return pass, nil
			}
		}
	}

	return "", ErrHashWasNotDecoded
}

func (i *Iterator) getChar(idx int) rune {
	return i.charList[idx]
}

func (i *Iterator) getFirstChar() rune {
	return i.charList[0]
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
