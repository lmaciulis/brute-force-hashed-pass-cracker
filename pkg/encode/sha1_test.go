package encode

import (
	"encoding/hex"
	"testing"
)

var (
	sha1InputStr   = "world"
	sha1OutputHash = "7c211433f02071597741e6ff5a8ea34789abbf43"
)

func TestEncoderSha1_Encode(t *testing.T) {
	e := NewEncoderSha1()

	hexSum, err := e.Encode([]byte(sha1InputStr))
	if err != nil {
		t.Errorf("Sha1 encoder result error not nil, got: %s, want: %s.", err.Error(), "nil")
	}

	resHash := hex.EncodeToString(hexSum)
	if resHash != sha1OutputHash {
		t.Errorf("Sha1 encoder result incorrect, got: %s, want: %s.", resHash, sha1OutputHash)
	}
}

func TestEncoderSha1_EncodeStr(t *testing.T) {
	e := NewEncoderSha1()
	res, err := e.EncodeStr(sha1InputStr)
	if err != nil {
		t.Errorf("Sha1 encoder result error not nil, got: %s, want: %s.", err.Error(), "nil")
	}
	if res != sha1OutputHash {
		t.Errorf("Sha1 encoder result incorrect, got: %s, want: %s.", res, sha1OutputHash)
	}
}

func TestEncoderSha1_Match(t *testing.T) {
	e := NewEncoderSha1()
	hexSum, _ := e.Encode([]byte(sha1InputStr))

	if !e.Match([]byte(sha1InputStr), hexSum) {
		t.Error("Sha1 encoder returned incorrect Match result")
	}
	if e.Match([]byte("not-match"), hexSum) {
		t.Error("Sha1 encoder returned incorrect Match result")
	}
}
