package encode

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
)

// EncoderSha1 sha1 encoder
type EncoderMd5 struct {
	Encoder hash.Hash
}

// EncodeStr encodes string to sha1 hash string
func (e *EncoderMd5) EncodeStr(pass string) (hash string, error error) {
	hexSum, _ := e.Encode([]byte(pass))

	return hex.EncodeToString(hexSum), nil
}

// Encode encodes bytes to md5 hash hex bytes
func (e *EncoderMd5) Encode(pass []byte) (hash []byte, err error) {
	e.Encoder.Reset()
	e.Encoder.Write(pass)

	return e.Encoder.Sum(nil), nil
}

// Match checks if given phrase (in bytes) matches hexadecimal hash (in bytes)
func (e *EncoderMd5) Match(pass []byte, hexHash []byte) bool {
	hexSum, _ := e.Encode(pass)

	return bytes.Equal(hexSum, hexHash)
}

// NewEncoderSha1 returns new EncoderSha1 instance
func NewEncoderMd5() *EncoderMd5 {
	return &EncoderMd5{
		Encoder: md5.New(),
	}
}
