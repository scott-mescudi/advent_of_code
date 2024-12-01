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
		l, _ := strconv.Atoi(stuff[0])
		w, _ := strconv.Atoi(stuff[1])
		h, _ := strconv.Atoi(stuff[2])

		xy := []int{(l * w), (w *h), (h * l)}
		slices.Sort(xy)

		x1 := 0
		for _, v := range xy {
			x1 += (2*v)
		}

		x1 += xy[0]
		
		total += x1
	}

	fmt.Println(total)
}