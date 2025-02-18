package main

import (
	"fmt"
	"math/rand/v2"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	height := 10
	width := 20
	canvas(height, width)
	return nil
}

func canvas(h, w int) {
	startX := rand.IntN(h)
	startY := rand.IntN(w)

	endX := rand.IntN(h)
	for endX == startX {
		endX = rand.IntN(h)
	}

	endY := rand.IntN(w)
	for endY == startY {
		endY = rand.IntN(w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			// better wall generation
			wallX := rand.IntN(3)
			wallY := rand.IntN(3)

			if i == startX && j == startY {
				fmt.Printf(White + "O" + Reset + " ")
			} else if i == endX && j == endY {
				fmt.Printf(White + "X" + Reset + " ")
			} else if i == wallX && j == wallY {

			} else {
				if wallX == 0 && wallY == 0 {
					fmt.Printf(Blue + "\u2587" + Reset + " ")
				} else {
					fmt.Printf(White + "\u2587" + Reset + " ")
				}
			}
		}
		fmt.Println()
	}
}
