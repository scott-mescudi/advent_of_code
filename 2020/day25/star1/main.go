package main

import "fmt"

func transform(value int, subjectNumber int) int {
	return (value * subjectNumber) % 20201227
}

func getLoopSize(key int) int {
	total := 0
	value := 1
	for {
		value2 := transform(value, 7)
		value = value2
		total++
		if key == value2 {
			break
		}
	}

	return total
}

func getKey(num, loopSize int) int {
	var num2 int = 1
	for i := 0; i < loopSize; i++ {
		num2 = transform(num2, num)
	}

	return num2
}

func main() {
	pk := 11404017
	dk := 13768789

	card := getLoopSize(pk)
	door := getLoopSize(dk)

	fmt.Println(getKey(dk, card))
	
	fmt.Println(card)
	fmt.Println(door)
}