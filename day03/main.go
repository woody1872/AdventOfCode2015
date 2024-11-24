package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"strings"
)

type Grid struct {
	x, y int
}

func (g *Grid) Up() {
	g.y += 1
}

func (g *Grid) Down() {
	g.y -= 1
}

func (g *Grid) Left() {
	g.x -= 1
}

func (g *Grid) Right() {
	g.x += 1
}

func (g Grid) String() string {
	return fmt.Sprintf("Grid{x:%d,y:%d}", g.x, g.y)
}

type GridVisitMap map[string]struct{}

func visit(gridVisitMap GridVisitMap, gridKey string) {
	_, exists := gridVisitMap[gridKey]
	if exists {
		gridVisitMap[gridKey] = struct{}{}
	} else {
		gridVisitMap[gridKey] = struct{}{}
	}
}

func main() {
	rawData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	directions := strings.Split(string(rawData), "")

	grid := Grid{0, 0}
	gridVisitMap := GridVisitMap{grid.String(): struct{}{}}
	for _, direction := range directions {
		switch direction {
		case "^":
			grid.Up()
			visit(gridVisitMap, grid.String())
		case "v":
			grid.Down()
			visit(gridVisitMap, grid.String())
		case "<":
			grid.Left()
			visit(gridVisitMap, grid.String())
		case ">":
			grid.Right()
			visit(gridVisitMap, grid.String())
		default:
			log.Fatal("unknown direction:", direction)
		}
	}
	fmt.Println("Answer (part 1):", len(gridVisitMap))

	santaGrid := Grid{0, 0}
	santaGridVisitMap := GridVisitMap{santaGrid.String(): struct{}{}}
	roboGrid := Grid{0, 0}
	roboGridVisitMap := GridVisitMap{roboGrid.String(): struct{}{}}
	for i, direction := range directions {
		switch direction {
		case "^":
			if i%2 == 0 {
				santaGrid.Up()
				visit(santaGridVisitMap, santaGrid.String())
			} else {
				roboGrid.Up()
				visit(roboGridVisitMap, roboGrid.String())
			}
		case "v":
			if i%2 == 0 {
				santaGrid.Down()
				visit(santaGridVisitMap, santaGrid.String())
			} else {
				roboGrid.Down()
				visit(roboGridVisitMap, roboGrid.String())
			}
		case "<":
			if i%2 == 0 {
				santaGrid.Left()
				visit(santaGridVisitMap, santaGrid.String())
			} else {
				roboGrid.Left()
				visit(roboGridVisitMap, roboGrid.String())
			}
		case ">":
			if i%2 == 0 {
				santaGrid.Right()
				visit(santaGridVisitMap, santaGrid.String())
			} else {
				roboGrid.Right()
				visit(roboGridVisitMap, roboGrid.String())
			}
		default:
			log.Fatal("unknown direction:", direction)
		}
	}
	combinedGridVisitMap := santaGridVisitMap
	maps.Copy(combinedGridVisitMap, roboGridVisitMap)
	fmt.Println("Answer (part 2):", len(combinedGridVisitMap))
}
