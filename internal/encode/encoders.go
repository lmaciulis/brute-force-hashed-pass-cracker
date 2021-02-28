package encode

const (
	algSha1 Algorithm = "sha1"
	algMd5  Algorithm = "md5"
)

type Algorithm string
type HashCode string
type PlainCode string

// Encoder interface for any type hash encoder
type Encoder interface {
	Encode(code PlainCode) (hash HashCode, err error)
}

// Factory returns appropriate encoder by provided algorithm
func Factory(alg Algorithm) (encoder Encoder, error error) {
	switch alg {
	case algSha1:
		return NewEncoderSha1(), nil
	case algMd5:
		return NewEncoderMd5(), nil
	default:
		return nil, ErrAlgorithmNotFound
	}
}
