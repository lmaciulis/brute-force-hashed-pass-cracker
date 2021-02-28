package encode

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

// EncoderSha1 sha1 encoder
type EncoderMd5 struct {
	Encoder hash.Hash
}

// Encode encodes string to sha1 hash string
func (e *EncoderMd5) Encode(code PlainCode) (hash HashCode, error error) {
	e.Encoder.Write([]byte(code))

	hexSum := e.Encoder.Sum(nil)
	strSum := hex.EncodeToString(hexSum)

	return HashCode(strSum), nil
}

// NewEncoderSha1 returns new EncoderSha1 instance
func NewEncoderMd5() *EncoderMd5 {
	return &EncoderMd5{
		Encoder: md5.New(),
	}
}
