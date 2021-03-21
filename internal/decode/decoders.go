package decode

const (
	AlgSha1 Alg = "sha1"
	AlgMd5  Alg = "md5"
)

type Alg string

// Decoder interface for any type hash decoder
type Decoder interface {
	Decode(hash string) (code string, err error)
}

// Factory returns appropriate decoder by provided algorithm
func Factory(alg Alg) (decoder Decoder, error error) {
	switch alg {
	case AlgSha1:
		return NewDecoderSha1(), nil
	case AlgMd5:
		return nil, ErrAlgorithmNotFound
		//return NewDecoderMd5(), nil
	default:
		return nil, ErrAlgorithmNotFound
	}
}
