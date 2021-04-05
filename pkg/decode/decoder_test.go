package decode

import (
	"testing"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/config"
)

func TestDecoder_Decode(t *testing.T) {
	decoder := NewDecoder("sha1", getCfg())
	input := "be9163bc410b18755091fa7a378f4c5fea2f09bf"
	exp := "1aba2"

	res, cnt, err := decoder.Decode(input)

	if err != nil {
		t.Errorf("decoder error not nil, got: %s, want: %s.", err.Error(), "nil")
	}
	if cnt < 50 {
		t.Errorf("decoder expected iterations count was too low, got: %d, want: at least %d.", cnt, 50)
	}
	if res != exp {
		t.Errorf("decoder result value incorrect, got: %s, want: %s.", res, exp)
	}
}

func TestDecoder_GetMaxIterations(t *testing.T) {
	decoder := NewDecoder("sha1", getCfg())
	maxIter := decoder.GetMaxIterations()

	if maxIter != 405 {
		t.Errorf("decoder expected max iterations incorect, got: %d, want: %d.", maxIter, 405)
	}
}

func TestDecoder_GetMaxIterationsCalcRepresent(t *testing.T) {
	decoder := NewDecoder("sha1", getCfg())
	msg := decoder.GetMaxIterationsCalcRepresent()
	exp := "3 (available chars)^3 (max phrase length) * 15 (suffix/prefix overhead) = 405 (max iterations)"

	if exp != msg {
		t.Errorf("decoder returned wrong calculations representation, got: %s, want: %s.", msg, exp)
	}
}

func getCfg() *config.Config {
	cfg := config.Config{
		AvailableChars: "abc",
		MaxPassLength:  3,
		Prefixes: config.PrefixesConfig{
			Enabled: true,
			List:    []string{"1", "2", "3"},
		},
		Suffixes: config.SuffixesConfig{
			Enabled: true,
			List:    []string{"1", "2", "3"},
		},
	}

	return &cfg
}
