package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: checksum [md5|sha1] <file name>")
		return
	}
	if os.Args[1] == "md5" {
		hash := md5.New()
		f, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		if _, err := io.Copy(hash, f); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%x\n", hash.Sum(nil))
		return
	} else if os.Args[1] == "sha1" {
		hash := sha1.New()
		f, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		if _, err := io.Copy(hash, f); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%x\n", hash.Sum(nil))
		return
	} else {
		fmt.Println("Usage: checksum [md5|sha1] <file name>")
		return
	}
}