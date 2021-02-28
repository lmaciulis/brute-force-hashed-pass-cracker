package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/encode"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal(encode.ErrToFewArgumentsPassed)
	}

	alg := encode.Algorithm(args[0])
	code := encode.PlainCode(args[1])
	encoder, err := encode.Factory(alg)

	if err != nil {
		log.Fatal(err)
	}

	hash, err := encoder.Encode(code)

	if err != nil {
		log.Fatal(err)
	}

	format := "For passcode '%s' and algorithm '%s' hash is:"

	fmt.Println(fmt.Sprintf(format, code, alg))
	fmt.Println(hash)
}
