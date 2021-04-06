package encode

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"hash"
)

// EncoderSha1 sha1 encoder
type EncoderSha1 struct {
	Encoder hash.Hash
}

// EncodeStr encodes string to sha1 hash string
func (e *EncoderSha1) EncodeStr(pass string) (hash string, error error) {
	hexSum, _ := e.Encode([]byte(pass))

	return hex.EncodeToString(hexSum), nil
}

// Encode encodes bytes to sha1 hash hex bytes
func (e *EncoderSha1) Encode(pass []byte) (hash []byte, err error) {
	e.Encoder.Reset()
	e.Encoder.Write(pass)

	return e.Encoder.Sum(nil), nil
}

// Match checks if given phrase (in bytes) matches hexadecimal hash (in bytes)
func (e *EncoderSha1) Match(pass []byte, hexHash []byte) bool {
	hexSum, _ := e.Encode(pass)

	return bytes.Equal(hexSum, hexHash)
}

// NewEncoderSha1 returns new EncoderSha1 instance
func NewEncoderSha1() *EncoderSha1 {
	return &EncoderSha1{
		Encoder: sha1.New(),
	}
}
