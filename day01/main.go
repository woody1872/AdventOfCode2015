package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	rawData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions := string(rawData)

	up := strings.Count(instructions, "(")
	down := strings.Count(instructions, ")")
	fmt.Println("Answer (part 1): ", up-down)

	pos := 0
	basement := 0
	for i, ins := range strings.Split(instructions, "") {
		if ins == ")" {
			pos -= 1
			if pos < 0 {
				basement = i + 1
				break
			}
		} else {
			pos += 1
		}
	}
	fmt.Println("Answer (part 2): ", basement)
}
