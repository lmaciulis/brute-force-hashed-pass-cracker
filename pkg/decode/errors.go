package decode

import "errors"

// Decode domain errors
var (
	ErrWrongArgumentsCount = errors.New("exactly 2 argument should be passed: hash algorithm and hash string to be decoded")
	ErrHashWasNotDecoded   = errors.New("unable to decode hash, iterations loop ended without success result")
)
