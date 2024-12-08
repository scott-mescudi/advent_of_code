package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

func validCross(x, y int, board [][]rune, dir rune) bool {
    switch dir {
    case '^':
        return y > 0 && board[y-1][x] != '#' && board[y-1][x] != 'X'
    case '>':
        return x < len(board[y])-1 && board[y][x+1] != '#' && board[y][x+1] != 'X'
    case 'v':
        return y < len(board)-1 && board[y+1][x] != '#' && board[y+1][x] != 'X'
    case '<':
        return x > 0 && board[y][x-1] != '#' && board[y][x-1] != 'X'
    }
    return false
}

func getBoard(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	board := [][]rune{}
	for scanner.Scan() {		
		board = append(board, []rune(scanner.Text()))
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

type pos struct {
	x,y int
}

type cross struct {
	x,y int
	dir rune
}

func movePlayere(board [][]rune, nx, ny int,np rune) ( []cross ){
	var (
		x   int = nx
		y   int = ny
		p   rune = np
		end bool = false
	)

	
	crosses := []cross{}
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

		if validCross(x, y, board, p) { 
   			crosses = append(crosses, cross{x, y, p})
		}
		
		if end {
			break
		}
	} 
	
	return crosses
}

func findStart(board [][]rune) (x,y int, dir rune) {
	for y1, row := range board {
		for x1, char := range row {
			if char == '^' || char == '>' || char == '<' || char == 'v' {
				return x1, y1, char
			}
		}
	}

	return 0, 0, 0
}

func lookForLoop(board [][]rune, nx, ny int,np rune) bool {
	var (
		x   int = nx
		y   int = ny
		p   rune = np
		end bool = false
	)
	visits := make([][]int, len(board))
	for i := range board {
		visits[i] = make([]int, len(board[i]))
	}

	visits[x][y]++
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

		visits[x][y]++

		if visits[x][y] >= 5 {
			return true
		}

		if end {
			break
		}
	} 

	return false
}

func CopyMatrix(matrix [][]rune) [][]rune {
	// Create a new matrix with the same dimensions
	rows := len(matrix)
	cols := len(matrix[0])
	newMatrix := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		newMatrix[i] = make([]rune, cols)
	}

	// Copy elements from the original matrix to the new matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			newMatrix[i][j] = matrix[i][j]
		}
	}
	return newMatrix
}

func main() {
	start := time.Now()
	board2 := getBoard("../data.txt")
	x, y, dir := findStart(board2)
    crosses := movePlayere(board2, x ,y, dir)

    var total int64 = 0
	var wg sync.WaitGroup
    for _, cross := range crosses {
		wg.Add(1)
		go func (){
			defer wg.Done()
			tb := CopyMatrix(board2)

			switch cross.dir {
			case '^':
				if cross.y != 0 {
					tb[cross.y-1][cross.x] = '#'
				}
			case '>':
				if cross.x != len(tb[cross.y])-1 {
					tb[cross.y][cross.x+1] = '#'
				}
			case 'v':
				if cross.y != len(tb)-1 {
					tb[cross.y+1][cross.x] = '#'
				}
			case '<':
				if cross.x != 0 {
					tb[cross.y][cross.x-1] = '#'
				}
			}

			if ok := lookForLoop(tb, x, y, dir); ok {
				atomic.AddInt64(&total, 1)
			}
		}()
    }

	wg.Wait()
    fmt.Println(total)
	fmt.Println(time.Since(start))
}