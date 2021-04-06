package decode

import (
	"encoding/hex"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/char"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/config"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/encode"
)

type Decoder struct {
	algorithm      encode.Alg
	charList       []rune
	charLen        int
	maxPassLen     int
	prefixes       [][]rune
	suffixes       [][]rune
	preEnabled     bool
	sufEnabled     bool
	preSufOverhead int
}

const (
	iterableRunesStarIndex = 1
	countOutputInterval    = 10
)

var (
	hash  []byte
	wg    sync.WaitGroup
	count int
)

// Decode main decoder method, iterates over all possible chars variations and checks with provided hash.
// Returns matched phrase or error if match not found, and all possible chars variations checked without success
func (i *Decoder) Decode(input string) (pass string, scanned int, err error) {
	hexHash, err := hex.DecodeString(input)

	if err != nil {
		return "", count, err
	}

	hash = hexHash
	holders := i.createHolders()
	ch := make(chan string, 1)
	chCnt := make(chan int)

	for _, h := range holders {
		go i.iterateHolder(h, ch, chCnt)
	}

	// closes all goroutines if program ended without success match
	go func() {
		wg.Wait()

		close(ch)
		close(chCnt)
	}()

	// outputs scanned phrases one every time interval
	go func() {
		ticker := time.NewTicker(countOutputInterval * time.Second)

		for {
			select {
			case <-ticker.C:
				fmt.Println(fmt.Sprintf("... phrases scanned: %d", count))
			}
		}
	}()

	// increments scanned phrases count
	go func() {
		for cnt := range chCnt {
			count += cnt
		}
	}()

	// returns phrase when success match found and sent to channel
	for res := range ch {
		return res, count, nil
	}

	return "", count, ErrHashWasNotDecoded
}

func (i *Decoder) iterateHolder(holder *char.Holder, ch chan string, chCnt chan int) {
	defer wg.Done()
	encoder, _ := encode.Factory(i.algorithm)

	for hIdx := iterableRunesStarIndex; hIdx < holder.GetLen(); hIdx++ {
		i.iterateHolderRune(holder, encoder, hIdx, ch, chCnt)
	}
}

func (i Decoder) iterateHolderRune(holder *char.Holder, encoder encode.Encoder, hIdx int, ch chan string, chCnt chan int) {
	isLastIteration := holder.GetLen() == hIdx+1

	for charIdx := 0; charIdx < i.charLen; charIdx++ {
		holder.Set(hIdx, i.getChar(charIdx))

		if encoder.Match(holder.ToBytes(), hash) {
			ch <- holder.ToString()
			return
		}

		chCnt <- 1

		if i.preEnabled || i.sufEnabled {
			wg.Add(1)
			go i.iteratePrefixesSuffixes(char.CloneHolder(holder), ch, chCnt)
		}

		if isLastIteration == false {
			i.iterateHolderRune(holder, encoder, hIdx+1, ch, chCnt)
		}
	}
}

func (i *Decoder) iteratePrefixesSuffixes(holder *char.Holder, ch chan string, chCnt chan int) {
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

	chCnt <- i.preSufOverhead
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

// NewDecoder creates new decoder by provided algorithm and decoder config
func NewDecoder(alg encode.Alg, cfg *config.Config) *Decoder {
	chars := []rune(cfg.AvailableChars)

	var prefixes [][]rune
	var suffixes [][]rune
	preSufOverhead := 0

	for _, c := range cfg.Prefixes.List {
		prefixes = append(prefixes, []rune(c))
	}
	for _, c := range cfg.Suffixes.List {
		suffixes = append(suffixes, []rune(c))
	}

	if cfg.Prefixes.Enabled {
		preSufOverhead += len(prefixes)
	}
	if cfg.Suffixes.Enabled {
		preSufOverhead += len(suffixes)
	}
	if cfg.Prefixes.Enabled && cfg.Suffixes.Enabled {
		preSufOverhead += len(prefixes) * len(suffixes)
	}

	return &Decoder{
		algorithm:      alg,
		charList:       chars,
		charLen:        len(chars),
		maxPassLen:     cfg.MaxPassLength,
		prefixes:       prefixes,
		suffixes:       suffixes,
		preEnabled:     cfg.Prefixes.Enabled,
		sufEnabled:     cfg.Suffixes.Enabled,
		preSufOverhead: preSufOverhead,
	}
}

// GetMaxIterations returns max available decoder iterations, calculated by provided config
func (i *Decoder) GetMaxIterations() int {
	base := math.Pow(float64(i.charLen), float64(i.maxPassLen))
	overhead := 1

	if i.preSufOverhead > 0 {
		overhead = i.preSufOverhead
	}

	return int(base * float64(overhead))
}

// GetMaxIterationsCalcRepresent returns max available decoder iterations, calculated by provided config,
// full explained mathematical string representation
func (i *Decoder) GetMaxIterationsCalcRepresent() string {
	overhead := 1
	if i.preSufOverhead > 0 {
		overhead = i.preSufOverhead
	}

	return fmt.Sprintf(
		"%d (available chars)^%d (max phrase length) * %d (suffix/prefix overhead) = %d (max iterations)",
		i.charLen,
		i.maxPassLen,
		overhead,
		i.GetMaxIterations())
}
