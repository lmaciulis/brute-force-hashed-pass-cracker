package decode

import (
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/char"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/encode"
)

type Iterator struct {
	encoder    encode.Encoder
	charList   []rune
	charLen    int
	maxPassLen int
}

const (
	iterableRunesStarIndex = 1
)

var (
	hash []byte
)

func (i *Iterator) Run(decHash []byte) (pass string, err error) {
	hash = decHash // set as global var

	//chFound := make(chan bool)
	//chRes := make(chan string)

	for _, h := range i.createHolders() {
		pass, err = i.iterateHolder(h)

		if err == nil {
			return pass, nil
		}
	}

	return "", ErrHashWasNotDecoded
}

func (i *Iterator) iterateHolder(holder *char.Holder) (pass string, err error) {
	// @todo implement channel
	// @todo get/add starting character postpone
	for hIdx := iterableRunesStarIndex; hIdx < holder.GetLen(); hIdx++ {
		pass, err = i.iterateHolderRune(holder, hIdx)

		if err == nil {
			return pass, nil
		}
	}

	return "", ErrHashWasNotDecoded
}

func (i Iterator) iterateHolderRune(holder *char.Holder, hIdx int) (pass string, err error) {
	isLastIteration := holder.GetLen() == hIdx+1

	for charIdx := 0; charIdx < i.charLen; charIdx++ {
		holder.Set(hIdx, i.getChar(charIdx))

		if i.encoder.Match(holder.ToBytes(), hash) {
			return holder.ToString(), nil
		}

		if isLastIteration == false {
			pass, err = i.iterateHolderRune(holder, hIdx+1)

			if err == nil {
				return pass, nil
			}
		}
	}

	return "", ErrHashWasNotDecoded
}

func (i *Iterator) createHolders() []*char.Holder {
	var holders []*char.Holder
	initChar := i.charList[0]

	// Creates all possible lengths char holders, with every possible starting character.
	// This allows efficiently split decoding jobs to separate go routines.
	// One characters holder per one go routine.
	for hLen := 1; hLen <= i.maxPassLen; hLen++ {
		for charIdx := 0; charIdx < i.charLen; charIdx++ {
			h := char.NewHolder(hLen, initChar)
			h.Set(0, i.getChar(charIdx))

			holders = append(holders, h)
		}
	}

	return holders
}

func (i *Iterator) getChar(idx int) rune {
	return i.charList[idx]
}

func NewIterator(encoder encode.Encoder) *Iterator {
	//chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	//	"abcdefghijklmnopqrstuvwxyz" +
	//	"0123456789")

	chars := []rune("abcdefghijklmnopqrstuvwxyz")

	return &Iterator{
		encoder:    encoder,
		charList:   chars,
		charLen:    len(chars),
		maxPassLen: 5,
	}
}
