package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getdata(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	var nums [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stuff := strings.Split(scanner.Text(), " ")
		cont := make([]int, 0)
		for _, v := range stuff {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}
			cont = append(cont, num)
		}

		nums = append(nums, cont)
	}

	return nums
}

func reverseSlice(slice []int) {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
}

func handleArr(input []int) bool {
	issafe := false
	arr := make([]int, len(input))
	copy(arr, input)
	size := len(arr)


	//temp = append(temp[:i], temp[i+1:]...)
	for i := 0; i < size; i++ {
		if i == size-1 {
			continue
		}
		if arr[i] - arr[i+1] >= 1 && arr[i] - arr[i+1] <= 3 {
			issafe = true
			continue
		}else {
			issafe = false
			break
		}
	}

		fmt.Println(input, arr, issafe)
	return issafe
}

func handleArrasc(input []int) bool {
	issafe := false
	arr := make([]int, len(input))
	copy(arr, input)
	size := len(arr)


	for i := 0; i < size; i++ {
		if i == size-1 {
			continue
		}
		if arr[i+1] - arr[i] >= 1 && arr[i+1] - arr[i] <= 3 {
			issafe = true
			continue
		}else {
			issafe = false
			break
		}
	}

	fmt.Println(input, arr, issafe)
	return issafe
}

func main() {
	data := getdata("../testdata.txt")
	total := 0
	unsafe := [][]int{}
	for _, arr := range data {
		var safe bool
		if arr[0] < arr [1] {
			safe = handleArrasc(arr)
		}else{
			safe = handleArr(arr)
		}

		if safe {total++}else { unsafe = append(unsafe, arr) }
	}

	fmt.Println(total)
	fmt.Println(len(unsafe))

}