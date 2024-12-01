package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)


func main() {
	file, err := os.Open("data.txt")
	if err != nil{
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)

	var sum int = 0
	for scanner.Scan() {
		var n1 string = ""
		var n2 string = ""
		for _, v := range scanner.Text(){
			if unicode.IsNumber(v) {
				if n1 == "" {
					n1 = string(v)
				}

				n2 = string(v)
			}
		}
		n1 += n2
		inty, _ := strconv.Atoi(n1)
		sum += inty
	}

	fmt.Println(sum)


}