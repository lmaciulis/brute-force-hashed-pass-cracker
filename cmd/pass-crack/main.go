package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/decode"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		log.Fatal(decode.ErrWrongArgumentsCount)
	}

	alg := decode.Alg(args[0])
	decoder, err := decode.Factory(alg)

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	pass, err := decoder.Decode(args[1])
	elapsed := time.Since(start)

	fmt.Println(fmt.Sprintf("Timing: %fs", elapsed.Seconds()))

	if err != nil {
		log.Fatal(err)
	}

	format := "Algorithm '%s', Pass '%s', Hash '%s'"
	fmt.Println(fmt.Sprintf(format, alg, pass, args[1]))
}
