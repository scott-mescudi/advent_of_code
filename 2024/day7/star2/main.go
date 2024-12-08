package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type target struct {
	TargetNum int
	Nums []int
}

func getData(filename string) ([]target, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rows := []target{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		var (
			targetstr string = strings.TrimSuffix(parts[0], ":")
			nums []int = []int{}
		)

		targetNum, err := strconv.Atoi(targetstr)
		if err != nil {
			return nil, err
		}

		for _, v := range parts[1:] {
			n, err :=  strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			nums = append(nums, n)
		}

		rows = append(rows, target{targetNum, nums})
	}

	return rows, nil
}

func concat(a, b int) int {
	multiplier := 1
	for temp := b; temp > 0; temp /= 10 {
		multiplier *= 10
	}

	return a*multiplier + b
}

func verifyRow(row *target) bool {
	target := row.TargetNum
	numbers := row.Nums[1:]
	reachable := []int{row.Nums[0]}

	for _, num := range numbers {
		newReach := []int{}
		for _, reach := range reachable {
			l1 := reach + num
			l2 := reach * num
			l3 := concat(reach, num)
			if l1 == target || l2 == target || l3 == target {
				return true
			}

			if l1 < target {
				newReach = append(newReach, l1)
			}

			if l2 < target {
				newReach = append(newReach, l2)
			}

			if l3 < target {
				newReach = append(newReach, l3)
			}
		}

		reachable = newReach
	}

	return false
}

func main() {
	data, err := getData("../data.txt")
	if err != nil {
		log.Println("Error:", err)
	}

	var sum int = 0
	for _, row := range data {
		if verifyRow(&row) {
			sum += row.TargetNum
		}
	}


	fmt.Println(sum)
}