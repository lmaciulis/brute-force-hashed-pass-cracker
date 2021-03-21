package decode

import (
	"encoding/hex"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/encode"
)

// DecoderSha1 sha1 decoder
type DecoderSha1 struct {
	Iterator *Iterator
}

func (d *DecoderSha1) Decode(hash string) (code string, err error) {
	hexDec, err := hex.DecodeString(hash)

	if err != nil {
		return "", err
	}

	pass, err := d.Iterator.Run(hexDec)

	if err != nil {
		return "", err
	}

	return pass, nil
}

func NewDecoderSha1() *DecoderSha1 {
	encoder := encode.NewEncoderSha1()

	return &DecoderSha1{
		Iterator: NewIterator(encoder),
	}
}
