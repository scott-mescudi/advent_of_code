package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)


func main() {
	file, err := os.Open("../data.txt")
	if err != nil{
		log.Fatalln(err)
	}

	scanner :=  bufio.NewScanner(file)

	total := 0
	for scanner.Scan(){
		stuff := strings.Split(scanner.Text(), "x")

		ribbon := 0
		bow := 0

		l, _ := strconv.Atoi(stuff[0])
		w, _ := strconv.Atoi(stuff[1])
		h, _ := strconv.Atoi(stuff[2])

		perms := []int{2 * (l+w), 2 * (w+h), 2 * (h +l)}
		slices.Sort(perms)
		
		ribbon += perms[0]

		bow += (l * w * h)

		fmt.Printf("Ribbon is %v, bow is %v, for a total of %v\n", ribbon, bow, ribbon+bow)
		total += ribbon + bow
	}

	fmt.Println(total)
}