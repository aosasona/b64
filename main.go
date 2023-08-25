package main

import (
	"b64/b64"
	"flag"
	"fmt"
	"log"
)

func main() {
	var shouldEncode, shouldDecode bool
	flag.BoolVar(&shouldEncode, "e", false, "encode a string")
	flag.BoolVar(&shouldDecode, "d", false, "decode a string")
	flag.Parse()

	data := flag.Arg(0)
	if data == "" {
		log.Fatal("Please provide a string to encode or decode")
	}

	// Make encoding the default operation
	if !shouldEncode && !shouldDecode {
		shouldEncode = true
	}

	var fn func(string) (string, error)

	if shouldEncode {
		fn = b64.Encode
	} else if shouldDecode {
		fn = b64.Decode
	}

	if fn == nil {
		log.Fatal("Please provide a valid flag")
	}

	result, err := b64.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("> " + result)
}
