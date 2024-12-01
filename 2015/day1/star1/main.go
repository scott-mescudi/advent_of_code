package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
	file, err := os.ReadFile("../data.txt")
	if err != nil{
		log.Fatalln(err)
	}
	s := string(file)

	count := 0
	for _, c := range s{
		if c == '('{
			count++
		}else if c == ')'{
			count--
		}
	} 

	fmt.Println(count)
}