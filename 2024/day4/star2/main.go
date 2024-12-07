package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(data [][]rune, yidx, xidx int) bool {
    if yidx < 1 || xidx < 1 || yidx >= len(data)-1 || xidx >= len(data[yidx])-1 {
        return false
    }

	left, right := false, false
	

	if data[yidx-1][xidx+1] == 'S' && data[yidx+1][xidx-1] == 'M' {
		right = true
	}

	if data[yidx-1][xidx-1] == 'S' && data[yidx+1][xidx+1] == 'M' {
		left = true
	}

	if data[yidx-1][xidx+1] == 'M' && data[yidx+1][xidx-1] == 'S' {
		right = true
	}

	if data[yidx-1][xidx-1] == 'M' && data[yidx+1][xidx+1] == 'S' {
		left = true
	}

	if right && left {
		return true
	}


	return false
}



func parseData(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	m1 := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		m2 := make([]rune, 0)
		for _, v := range txt {
			m2 = append(m2, v)
		}

		m1 = append(m1, m2)
	}

	return m1
}

func main() {

	m1 := parseData("../data.txt")
	total := 0

	for yindex, l := range m1 {
		for xindex, r := range l {
			if r == 65 {
				if ok := check(m1, yindex, xindex); ok {
					total++
				}
			}	
		}
	}


	fmt.Println(total)
}