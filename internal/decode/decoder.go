package decode

import (
	"encoding/hex"
	"sync"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/char"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/config"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/encode"
)

type Decoder struct {
	algorithm  encode.Alg
	charList   []rune
	charLen    int
	maxPassLen int
	prefixes   [][]rune
	suffixes   [][]rune
	preEnabled bool
	sufEnabled bool
}

const (
	iterableRunesStarIndex = 1
)

var (
	hash []byte
	wg   sync.WaitGroup
)

func (i *Decoder) Decode(input string) (pass string, err error) {
	hexHash, err := hex.DecodeString(input)

	if err != nil {
		return "", err
	}

	hash = hexHash
	holders := i.createHolders()
	ch := make(chan string, len(holders)) // @todo remove buffer when routines wil be added dynamiclly

	for _, h := range holders {
		go i.iterateHolder(h, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		return res, nil
	}

	return "", ErrHashWasNotDecoded
}

func (i *Decoder) iterateHolder(holder *char.Holder, ch chan string) {
	defer wg.Done()
	encoder, _ := encode.Factory(i.algorithm)

	for hIdx := iterableRunesStarIndex; hIdx < holder.GetLen(); hIdx++ {
		if pass := i.iterateHolderRune(holder, encoder, hIdx); pass != "" {
			ch <- pass
		}
	}
}

func (i Decoder) iterateHolderRune(holder *char.Holder, encoder encode.Encoder, hIdx int) string {
	isLastIteration := holder.GetLen() == hIdx+1

	for charIdx := 0; charIdx < i.charLen; charIdx++ {
		holder.Set(hIdx, i.getChar(charIdx))
		// @todo create separate holder routines for prefixes and suffixes

		if encoder.Match(holder.ToBytes(), hash) {
			return holder.ToString()
		}

		if isLastIteration == false {
			if pass := i.iterateHolderRune(holder, encoder, hIdx+1); pass != "" {
				return pass
			}
		}
	}

	return ""
}

func (i *Decoder) createHolders() []*char.Holder {
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
			wg.Add(1)
		}
	}

	return holders
}

func (i *Decoder) getChar(idx int) rune {
	return i.charList[idx]
}

func NewDecoder(alg encode.Alg, cfg *config.Config) *Decoder {
	chars := []rune(cfg.AvailableChars)

	var prefixes [][]rune
	var suffixes [][]rune

	for _, c := range cfg.Prefixes.List {
		prefixes = append(prefixes, []rune(c))
	}
	for _, c := range cfg.Suffixes.List {
		suffixes = append(suffixes, []rune(c))
	}

	return &Decoder{
		algorithm:  alg,
		charList:   chars,
		charLen:    len(chars),
		maxPassLen: cfg.MaxPassLength,
		prefixes:   prefixes,
		suffixes:   suffixes,
		preEnabled: cfg.Prefixes.Enabled,
		sufEnabled: cfg.Suffixes.Enabled,
	}
}
