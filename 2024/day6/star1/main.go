package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func getBoard(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	board := [][]rune{}
	for scanner.Scan() {
		row := []rune{}
		for _, v := range scanner.Text() {
			row = append(row, v)
		}
		board = append(board, row)
	}

	return board
}

func moveUp(board [][]rune, x, y int) (newx int, newy int, end bool, dirchange rune) {
	if y == 0  {
		return x, y, true, '^'
	}

	if board[y-1][x] == '#' {
		return x, y, false, '>'
	}

	board[y][x] = 'X'
	board[y-1][x] = '^'
	return x, y - 1, false, '^'
}

func moveDown(board [][]rune, x, y int) (newx int, newy int, end bool, dirchange rune) {
	if y == len(board)-1 {
		return x, y, true, 'v'
	}

	if board[y+1][x] == '#' {
		return x, y, false, '<'	
	}

	board[y][x] = 'X'
	board[y+1][x] = 'v'
	return x, y + 1, false, 'v'
}

func moveRight(board [][]rune, x, y int) (newx int, newy int, end bool, dirchange rune) {
	if x == len(board[y])-1 {
		return x, y, true, '>'
	}

	if board[y][x+1] == '#' {
		return x, y, false, 'v'
	}

	board[y][x] = 'X'
	board[y][x+1] = '>'
	return x+1, y, false, '>'
}

func moveLeft(board [][]rune, x, y int) (newx int, newy int, end bool, dirchange rune) {
	if x == 0 {
		return x, y, true, '<'
	}

	if board[y][x-1] == '#' {
		return x, y, false, '^'
	}

	board[y][x] = 'X'
	board[y][x-1] = '<'
	return x-1, y, false, '<'
}

func printBoard(board [][]rune) {
	for _, row := range board {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func countSteps(board [][]rune) int {
	total := 0
	for _, row := range board {
		for _, char := range row {
			if char == 'X'{
				total++
			}
		}
	}

	return total+1
}

func main() {
	start :=  time.Now()
	board := getBoard("../data.txt")

	var (
		x   int
		y   int
		p   rune
		end bool
	)
	for y1, row := range board {
		for x1, char := range row {
			if char == '^' || char == '>' || char == '<' || char == 'v' {
				x = x1
				y = y1
				p = char
				end = false
			}
		}
	}



	for {
		switch p {
		case '^':
			nx, ny, e, dc := moveUp(board, x, y)
			x, y, p, end = nx, ny, dc, e
		case '>':
			nx, ny, e, dc := moveRight(board, x, y)
			x, y, p, end = nx, ny, dc, e
		case 'v':
			nx, ny, e, dc := moveDown(board, x, y)
			x, y, p, end = nx, ny, dc, e
		case '<':
			nx, ny, e, dc := moveLeft(board, x, y)
			x, y, p, end = nx, ny, dc, e
		}

		if end {
			break
		}
	}

	
	fmt.Println(countSteps(board))
	fmt.Println(time.Since(start))
}
