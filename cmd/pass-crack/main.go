package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/decode"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/internal/encode"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		log.Fatal(decode.ErrWrongArgumentsCount)
	}

	alg := encode.Alg(args[0])
	decoder := decode.NewDecoder(alg)

	// add for timing https://tour.golang.org/concurrency/6
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
