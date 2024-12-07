package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func horizontal(data []rune, xindex int) bool {
	if xindex > len(data)-4 {
		return false
	}

	if data[xindex+1] == 'M' && data[xindex+2] == 'A' && data[xindex+3] == 'S'{
		return true
	}

	return false
}

func reverse(data []rune, xindex int) bool {
	if xindex < 3 {
		return false
	}

	if data[xindex-1] == 'M' && data[xindex-2] == 'A' && data[xindex-3] == 'S'{
		return true
	}

	return false
}
func verticalup(data [][]rune, yindex, xindex int) bool {
	if yindex < 3 {
		return false
	}


	if data[yindex-1][xindex] == 'M' && data[yindex-2][xindex] == 'A' && data[yindex-3][xindex] == 'S' {
		return true
	}

	return false
}

func verticalDown(data [][]rune, yindex, xindex int) bool {
	if yindex > len(data)-4{
		return false
	}


	if data[yindex+1][xindex] == 'M' && data[yindex+2][xindex] == 'A' && data[yindex+3][xindex] == 'S' {
		return true
	}

	return false
}

func diagonaltopright(data [][]rune, yindex, xindex int) bool {
	if yindex < 3 || xindex > len(data[yindex])-4 {
		return false
	}
	
	if data[yindex-1][xindex+1] == 'M' && data[yindex-2][xindex+2] == 'A' && data[yindex-3][xindex+3] == 'S'{
		return true
	}

	return false
}

func diagonaltopleft(data [][]rune, yindex, xindex int) bool {
	if  yindex < 3 || xindex < 3 {
		return false
	}
	
	if data[yindex-1][xindex-1] == 'M' && data[yindex-2][xindex-2] == 'A' && data[yindex-3][xindex-3] == 'S'{
		return true
	}

	return false
}


func diagonalbottomleft(data [][]rune, yindex, xindex int) bool {
	if yindex > len(data) - 4 || xindex < 3 {
		return false
	}

	if data[yindex+1][xindex-1] == 'M' && data[yindex+2][xindex-2] == 'A' && data[yindex+3][xindex-3] == 'S' {
		return true
	}


	return false
}

func diagonalbottomright(data [][]rune, yindex, xindex int) bool {
	if yindex > len(data) - 4 || xindex > len(data[yindex])-4 {
		return false
	}

	if data[yindex+1][xindex+1] == 'M' && data[yindex+2][xindex+2] == 'A' && data[yindex+3][xindex+3] == 'S' {
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
			if r == 88 {
				if ok := horizontal(l, xindex); ok {
					total++
				}

				if ok := reverse(l, xindex); ok {
					total++
				}

				if ok := verticalup(m1, yindex, xindex); ok{
					total++
				}

				
				if ok := verticalDown(m1, yindex, xindex); ok{
					total++
				}

				if ok := diagonaltopleft(m1, yindex, xindex); ok{
					total++
				}

				if ok := diagonaltopright(m1, yindex, xindex); ok{
					total++
				}

				if ok := diagonalbottomright(m1, yindex, xindex); ok{
					total++
				}
				
				if ok := diagonalbottomleft(m1, yindex, xindex); ok{
					total++
				}
			}	
		}
	}

	fmt.Println(total)
}