package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	data := "bgvyzdsv"
	answer := ""
	i := 0
	for answer == "" {
		hash := md5.Sum([]byte(data + strconv.Itoa(i)))
		if hex.EncodeToString(hash[:])[0:6] == "000000" {
			answer = hex.EncodeToString(hash[:])
			fmt.Println(answer)
			fmt.Println(i)
			break
		}
		i++

	}

}
