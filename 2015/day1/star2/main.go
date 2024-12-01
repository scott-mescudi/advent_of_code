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
	for i, c := range s{
		if count == -1 {
			fmt.Println(i)
			break
		}

		if c == '('{
			count++
		}else if c == ')'{
			count--
		}
	} 

	fmt.Println(count)
}