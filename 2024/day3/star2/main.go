package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getok(data string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	total := 0

	for _, v := range matches {
		n1, err := strconv.Atoi(v[1])
		n2, err := strconv.Atoi(v[2])
		if err != nil {
			log.Fatalln(err)
		}

		total += n1 * n2
	}

	return total
}

//do()mul(2,3)don't()mul(4,5)do()mul(6,7)


func main() {
	data, err := os.ReadFile("../testdata.txt")
	if err != nil {
		log.Fatalln(err)
	}
	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	matches := re.FindStringSubmatch(string(data))

	stuff := string(data)


	for _, v := range matches {
		stuff = strings.ReplaceAll(stuff, v, " ")
	}


	fmt.Println(stuff)

	fmt.Println(getok(string(stuff)))
}
