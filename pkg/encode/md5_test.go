package encode

import (
	"encoding/hex"
	"testing"
)

var (
	md5InputStr   = "hello"
	md5OutputHash = "5d41402abc4b2a76b9719d911017c592"
)

func TestEncoderMd5_Encode(t *testing.T) {
	e := NewEncoderMd5()

	hexSum, err := e.Encode([]byte(md5InputStr))
	if err != nil {
		t.Errorf("MD5 encoder result error not nil, got: %s, want: %s.", err.Error(), "nil")
	}

	resHash := hex.EncodeToString(hexSum)
	if resHash != md5OutputHash {
		t.Errorf("MD5 encoder result incorrect, got: %s, want: %s.", resHash, md5OutputHash)
	}
}

func TestEncoderMd5_EncodeStr(t *testing.T) {
	e := NewEncoderMd5()
	res, err := e.EncodeStr(md5InputStr)
	if err != nil {
		t.Errorf("MD5 encoder result error not nil, got: %s, want: %s.", err.Error(), "nil")
	}
	if res != md5OutputHash {
		t.Errorf("MD5 encoder result incorrect, got: %s, want: %s.", res, md5OutputHash)
	}
}

func TestEncoderMd5_Match(t *testing.T) {
	e := NewEncoderMd5()
	hexSum, _ := e.Encode([]byte(md5InputStr))

	if !e.Match([]byte(md5InputStr), hexSum) {
		t.Error("MD5 encoder returned incorrect Match result")
	}
	if e.Match([]byte("not-match"), hexSum) {
		t.Error("MD5 encoder returned incorrect Match result")
	}
}
