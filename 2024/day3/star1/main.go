package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("../testdata.txt")
	if err != nil {
		log.Fatalln(err)
	}
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(string(data), -1)

	total := 0

	for _, v := range matches {
		n1, err := strconv.Atoi(v[1])
		n2, err := strconv.Atoi(v[2])
		if err != nil {
			log.Fatalln(err)
		}

		total += n1 * n2
	}

	fmt.Println(total)
}
// `don’t\(\)(.*?)do\(\)`
