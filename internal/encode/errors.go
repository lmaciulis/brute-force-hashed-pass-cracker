package encode

import "errors"

// Encode domain errors
var (
	ErrWrongArgumentsCount = errors.New("exactly 2 argument should be passed: hash algorithm and string to be hashed")
	ErrAlgorithmNotFound   = errors.New("such hash algorithm is not configured")
)
