package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)
	floor := 0

	firstBaseFound := false
	for i := 0; i < len(input); i++ {
		if input[i] == 40 {
			floor++
		} else {
			floor--
		}
		if floor == -1 && firstBaseFound == false {
			firstBaseFound = true
			fmt.Println("Basement = " + strconv.Itoa(i+1))
		}
	}
	fmt.Println(floor)
}
