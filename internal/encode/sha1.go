package encode

import (
	"crypto/sha1"
	"encoding/hex"
	"hash"
)

// EncoderSha1 sha1 encoder
type EncoderSha1 struct {
	Encoder hash.Hash
}

// Encode encodes string to sha1 hash string
func (e *EncoderSha1) Encode(code PlainCode) (hash HashCode, error error) {
	e.Encoder.Write([]byte(code))

	hexSum := e.Encoder.Sum(nil)
	strSum := hex.EncodeToString(hexSum)

	return HashCode(strSum), nil
}

// NewEncoderSha1 returns new EncoderSha1 instance
func NewEncoderSha1() *EncoderSha1 {
	return &EncoderSha1{
		Encoder: sha1.New(),
	}
}
