package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hashString = flag.String("hash-string", "sha256", "Write string")
var hashAlgo = flag.Int("hash-algo", 256, "Select preffered alghoritm(sha256, sha384, sha512)")

func main() {
	flag.Parse()

	input := *hashString

	switch *hashAlgo {
	case 384:
		fmt.Printf("sha384 for %q: %x\n", input, sha512.Sum384([]byte(input)))
	case 512:
		fmt.Printf("sha512 for %q: %x\n", input, sha512.Sum512([]byte(input)))
	default:
		fmt.Printf("sha512 for %q: %x\n", input, sha256.Sum256([]byte(input)))
	}
}
