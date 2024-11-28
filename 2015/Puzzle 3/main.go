package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func listInit() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	str := string(data)
	return str
}
func move(input string, pos coordinates) coordinates {

	if input == "v" {
		pos.y -= 1
	}
	if input == "^" {
		pos.y += 1
	}
	if input == "<" {
		pos.x -= 1
	}
	if input == ">" {
		pos.x += 1
	}
	houseCoordinates := coordinates{pos.x, pos.y}
	var isUnique bool = true
	for _, i := range coordinateList {
		if houseCoordinates.x == i.x && houseCoordinates.y == i.y {
			isUnique = false
			break
		}
	}
	if isUnique {
		unique = append(unique, houseCoordinates)
	}
	return houseCoordinates
}

type coordinates struct {
	x, y int
}

var fleshCoordinates coordinates = coordinates{0, 0}
var roboCoordinates coordinates = coordinates{0, 0}
var coordinateList []coordinates
var unique []coordinates

func moveFlesh(input string) {
	fleshCoordinates = move(input, fleshCoordinates)
	coordinateList = append(coordinateList, fleshCoordinates)
}
func moveRobo(input string) {
	roboCoordinates = move(input, roboCoordinates)
	coordinateList = append(coordinateList, roboCoordinates)
}
func main() {
	// unique = append(unique, coordinates{0, 0})
	directions := string(listInit())
	var isFleshy bool = true
	//fmt.Println(directions)

	for _, i := range directions {
		input := string(i)
		if isFleshy {
			moveFlesh(input)
			isFleshy = false
		} else if !isFleshy {
			moveRobo(input)
			isFleshy = true
		}
	}
	//fmt.Println(coordinateList)
	fmt.Println(len(unique))

}
