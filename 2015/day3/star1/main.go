package main

import (
	"fmt"
	"log"
	"os"
)

type house struct {
	n,e,s,w int
}

type santa struct {
	n,e,s,w int	
}

func main() {
	file, err := os.ReadFile("../data.txt")
	if err != nil{
		log.Fatalln(err)
	}
	

	santa := santa{0, 0, 0,0}
	locations := map[house]int{}
	locations[house(santa)]++

	for _, c  := range file {
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
	fmt.Println(len(locations))
}