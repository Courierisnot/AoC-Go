package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	data := "abcdef609043"
	hash := md5.Sum([]byte(data))
	fmt.Println(hex.EncodeToString(hash[:]))
}
