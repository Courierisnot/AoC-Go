package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func listInit() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return lines

}
func getSmallestSideAndPerim(list []int) (int, int) {
	sort.Ints(list)
	side := list[0] * list[1]
	smolPerim := (list[0] * 2) + (list[1] * 2)
	return side, smolPerim

}
func getSurfaceArea(l int, w int, h int) int {
	surfaceArea := ((l * w) + (l * h) + (w * h)) * 2
	return surfaceArea
}
func convertToNum(box string) (int, int, int) {
	splitList := strings.Split(box, "x")
	l, err := strconv.Atoi(splitList[0])
	if err != nil {
		log.Fatal()
	}
	w, err := strconv.Atoi(splitList[1])
	if err != nil {
		log.Fatal()
	}
	h, err := strconv.Atoi(splitList[2])
	if err != nil {
		log.Fatal()
	}

	return l, w, h
}
func main() {

	list := listInit()
	total := 0
	ribbon := 0
	for _, box := range list {
		l, w, h := convertToNum(box)
		area := getSurfaceArea(l, w, h)
		smolSide, smolPerim := getSmallestSideAndPerim([]int{l, w, h})
		ribbon += smolPerim + (l * w * h)
		total += area + smolSide
	}
	fmt.Println(total)
	fmt.Println(ribbon)

}
