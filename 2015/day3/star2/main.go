package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type house struct {
	n,e,s,w int
}

type Santa struct {
	n,e,s,w int	
}

func main() {
	start := time.Now()
	file, err := os.ReadFile("../data.txt")
	if err != nil{
		log.Fatalln(err)
	}

	
	var santaRoute strings.Builder
	var roboRoute strings.Builder

	sturn := false
	for _, c := range file {
		if !sturn{
			santaRoute.WriteByte(c)
			sturn = true
			continue
		}

		if sturn {
			roboRoute.WriteByte(c)
			sturn = false
			continue
		}
	}


	santa := Santa{0, 0, 0,0}
	roboSanta := Santa{0, 0, 0,0}

	locations := map[house]int{}
	locations[house(santa)]++

	

	for _, c  := range santaRoute.String() {
		switch c {
		case '^':
			santa.n++
		case 'v':
			santa.n--
		case '<':
			santa.w--
		case '>':
			santa.w++
		}

		locations[house(santa)]++
	}

	for _, c  := range roboRoute.String() {
		switch c {
		case '^':
			roboSanta.n++
		case 'v':
			roboSanta.n--
		case '<':
			roboSanta.w--
		case '>':
			roboSanta.w++
		}

		locations[house(roboSanta)]++
	}

	fmt.Println(len(locations))
	fmt.Println(time.Since(start))
}