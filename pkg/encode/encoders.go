package encode

const (
	AlgSha1 Alg = "sha1"
	AlgMd5  Alg = "md5"
)

type Alg string

// Encoder interface for any type hash encoder
type Encoder interface {
	EncodeStr(pass string) (hash string, err error)
	Encode(pass []byte) (hash []byte, err error)
	Match(pass []byte, hexHash []byte) bool
}

// Factory returns appropriate encoder by provided algorithm
func Factory(alg Alg) (encoder Encoder, error error) {
	switch alg {
	case AlgSha1:
		return NewEncoderSha1(), nil
	case AlgMd5:
		return NewEncoderMd5(), nil
	default:
		return nil, ErrAlgorithmNotFound
	}
}
