package main

import (
	"fmt"
	"encoding/base64"
	"os"
)

func main () {
	if len(os.Args) != 3 {
		Usage()
		return
	}

	action := os.Args[1]
	str := os.Args[2]
	if action == "-e" || action == "-E" {
		fmt.Printf(base64.RawURLEncoding.EncodeToString([]byte(str)))
	} else if action == "-d" || action == "-D" {
		res, err := base64.RawURLEncoding.DecodeString(str)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%q", res)
	}
	
}

func Usage() {
	fmt.Println("Usage: gobase64 [-e encoding |-d decoding] string")
}