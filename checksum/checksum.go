package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/wenchma/goutils"
)

func main() {
	if len(os.Args) != 3 {
		Usage()
		return
	}
	v, err := CalculateSum(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	table := goutils.NewTable([]string{"Algorithm", "Value"})
	table.AddRow(map[string]interface{}{"Algorithm": "sha1", "Value": v})
	table.Print()
	return
}

func Usage() {
	fmt.Println("Usage: checksum [md5|sha1] <file name>")
}

func CalculateSum(algorithm, file string) (string, error) {
	NewFuncs := map[string]func() hash.Hash {
		"md5": md5.New,
		"sha1": sha1.New,
	}
	NewFunc, ok := NewFuncs[algorithm]
	if !ok {
		return "", fmt.Errorf("Unsupported Hash Algorithm: %s", algorithm)
	}
	hash := NewFunc()
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}