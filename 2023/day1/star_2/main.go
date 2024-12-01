package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"unicode"
)

type shi struct {
	v string
	index int
}

type shiz []shi

func (a shiz) Len() int           { return len(a) }
func (a shiz) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a shiz) Less(i, j int) bool { return a[i].index < a[j].index }

func main() {
	file, err := os.Open("data.txt")
	if err != nil{
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)

	var sum int = 0

	for scanner.Scan() {
		stuff := []shi{}

		ll := scanner.Text()

		for i, v := range ll{
			if unicode.IsNumber(v) {
				stuff = append(stuff, shi{v: string(v), index: i})
			}
		}

		numbers := []string{
			"one",
			"two",
			"three",
			"four",
			"five",
			"six",
			"seven",
			"eight",
			"nine",
		}

		m := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}


		for _, v := range numbers {
			index := strings.Index(strings.ToLower(ll), v)
			if index == -1 {
				continue
			}
			val, _ := m[v]
			stuff = append(stuff, shi{v: val, index: index})
		}

		sort.Sort(shiz(stuff))

		if len(stuff) == 1{
			stuff = append(stuff, stuff[0])
		}

		var n3 string
		size := len(stuff)
		n3 = stuff[0].v + stuff[size-1].v
		
		inty, err := strconv.Atoi(n3)
		if err != nil {
			log.Fatalln(err)
		}

		sum += inty
	}

	fmt.Println(sum)
}