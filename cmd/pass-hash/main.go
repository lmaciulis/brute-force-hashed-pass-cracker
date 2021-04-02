package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/encode"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		log.Fatal(encode.ErrWrongArgumentsCount)
	}

	alg := encode.Alg(args[0])
	encoder, err := encode.Factory(alg)

	if err != nil {
		log.Fatal(err)
	}

	hash, err := encoder.EncodeStr(args[1])

	if err != nil {
		log.Fatal(err)
	}

	format := "Algorithm '%s', Pass '%s', Hash '%s'"
	fmt.Println(fmt.Sprintf(format, alg, args[1], hash))
}
