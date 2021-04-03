package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/config"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/decode"
	"github.com/lmaciulis/brute-force-hashed-pass-cracker/pkg/encode"
)

var cfg config.Config

func init() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		log.Fatal(decode.ErrWrongArgumentsCount)
	}

	alg := encode.Alg(args[0])
	decoder := decode.NewDecoder(alg, &cfg)

	printIntro(decoder, alg, args[1])

	start := time.Now()
	pass, scanned, err := decoder.Decode(args[1])
	elapsed := time.Since(start)

	fmt.Println(fmt.Sprintf("Time elapsed: %f mins", elapsed.Minutes()))
	fmt.Println(fmt.Sprintf("Total phrases scanned: %d", scanned))

	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s, Hash '%s'", err.Error(), args[1]))
	} else {
		fmt.Println(fmt.Sprintf("Success: algorithm '%s', Pass '%s', Hash '%s'", alg, pass, args[1]))
	}
}

func printIntro(decoder *decode.Decoder, alg encode.Alg, hash string) {
	format := `--- Brute Force Hashed Phrase Decoder ---
You are trying to decode hash '%s' with '%s' algorithm.
Please be aware, that every additional character added in config exponentially increases program execution time.
As an example:
Standard PC, with 4 CPU cores ~2.5 Ghz, scans around 1 million phrases per seconds for SHA1 algorithm (MD5 works much faster).
So, for SHA1 algorithm, based on your current configuration there are:
%s
As a result max execution time can lead up to: %f hours.
You can always terminate program by pressing 'ctrl/cmd + c'

Starting program....
`

	fmt.Println(fmt.Sprintf(
		format,
		hash,
		alg,
		decoder.GetMaxIterationsCalcRepresent(),
		float64(decoder.GetMaxIterations())/1000000/60/60))
}
