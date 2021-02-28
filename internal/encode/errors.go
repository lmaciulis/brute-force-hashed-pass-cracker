package encode

import "errors"

// Domain errors.
var (
	ErrToFewArgumentsPassed = errors.New("at least 2 argument should be passed: hash algorithm and at least one string to be hashed")
	ErrAlgorithmNotFound    = errors.New("such hash algorithm is not configured")
)
