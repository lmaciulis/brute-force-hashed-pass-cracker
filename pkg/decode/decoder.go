package decode

import (
	"encoding/hex"
	"sync"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/char"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/config"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/encode"
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
	ch := make(chan string, 1)

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
		i.iterateHolderRune(holder, encoder, hIdx, ch)
	}
}

func (i Decoder) iterateHolderRune(holder *char.Holder, encoder encode.Encoder, hIdx int, ch chan string) {
	isLastIteration := holder.GetLen() == hIdx+1

	for charIdx := 0; charIdx < i.charLen; charIdx++ {
		holder.Set(hIdx, i.getChar(charIdx))

		if encoder.Match(holder.ToBytes(), hash) {
			ch <- holder.ToString()
			return
		}

		if i.preEnabled || i.sufEnabled {
			wg.Add(1)
			go i.iteratePrefixesSuffixes(char.CloneHolder(holder), ch)
		}

		if isLastIteration == false {
			i.iterateHolderRune(holder, encoder, hIdx+1, ch)
		}
	}
}

func (i *Decoder) iteratePrefixesSuffixes(holder *char.Holder, ch chan string) {
	defer wg.Done()

	encoder, _ := encode.Factory(i.algorithm)

	if i.preEnabled {
		// loop through prepended prefixes and check if hash match
		for _, prefix := range i.prefixes {
			hp := char.CloneHolderWithPrefix(holder, prefix)
			if encoder.Match(hp.ToBytes(), hash) {
				ch <- hp.ToString()
				return
			}

			if i.sufEnabled {
				// also append suffix for each prefixed handler and check if hash match
				for _, suffix := range i.suffixes {
					hs := char.CloneHolderWithSuffix(hp, suffix)
					if encoder.Match(hs.ToBytes(), hash) {
						ch <- hs.ToString()
						return
					}
				}
			}
		}
	}

	if i.sufEnabled {
		// finally, loop through original passed handler with only available suffixes appended, and check if hash match
		for _, suffix := range i.suffixes {
			hs := char.CloneHolderWithSuffix(holder, suffix)
			if encoder.Match(hs.ToBytes(), hash) {
				ch <- hs.ToString()
				return
			}
		}
	}
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
