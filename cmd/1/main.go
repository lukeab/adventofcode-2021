package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2021/pkg/config"
	"github.com/lukeab/adventofcode-2021/pkg/fileloader"
)

func main() {
	fmt.Println("Advent of code 2021 Day 1")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/1/input"

	intslice, err := fileloader.LoadFileLinesAsIntSlice(cfg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:")

	var bcounter, i, x, y int
	fmt.Printf("Length of list: %d\n", len(intslice))
	for i = 0; i < len(intslice)-1; i++ {
		x = intslice[i]
		y = intslice[i+1]
		fmt.Printf("%d > %d", y, x)
		if y > x {
			bcounter++
			fmt.Print(" - y")
		}
		fmt.Print("\n")
	}
	fmt.Println(bcounter)
}
