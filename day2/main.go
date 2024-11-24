package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Present struct {
	length int
	width  int
	height int
}

func (p Present) surfaceArea() int {
	return (2 * (p.length * p.width)) + (2 * (p.width * p.height)) + (2 * (p.height * p.length))
}

func (p Present) WrappingPaper() int {
	area := p.surfaceArea()
	slack := min(p.length*p.width, p.width*p.height, p.height*p.length)
	return area + slack
}

func (p Present) Ribbon() int {
	wrap := 2 * min(p.length+p.width, p.width+p.height, p.height+p.length)
	bow := p.length * p.width * p.height
	return wrap + bow
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	var presents []Present
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "x")
		l, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		h, _ := strconv.Atoi(parts[2])
		present := Present{length: l, width: w, height: h}
		presents = append(presents, present)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	orderWrapping := 0
	orderRibbon := 0
	for _, p := range presents {
		orderWrapping += p.WrappingPaper()
		orderRibbon += p.Ribbon()
	}
	fmt.Println("Answer (part 1):", orderWrapping)
	fmt.Println("Answer (part 2):", orderRibbon)
}
