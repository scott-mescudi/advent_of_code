package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getok(data string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	total := 0

	for _, v := range matches {
		n1, err := strconv.Atoi(v[1])
		n2, err := strconv.Atoi(v[2])
		if err != nil {
			log.Fatalln(err)
		}

		total += n1 * n2
	}

	return total
}




type indices struct {
	idx, lidx int
}


func main() {
	data, err := os.ReadFile("../data.txt")
	if err != nil {
		log.Fatalln(err)
	}	

	const (
		w1 = "don't()"
		w2 = "do()"
	)

	ids := []indices{}

	var lidx int = 0
	for {
		idx := strings.Index(string(data[lidx:]), w1)
		if idx == -1 {
			break
		}
		idx += lidx 

		
		idx2 := strings.Index(string(data[idx+len(w1):]), w2)
		if idx2 == -1 {
			oc1 := indices{idx: idx, lidx: len(data)}
			ids = append(ids, oc1)

			break
		}

		idx2 += idx + len(w1) 
		oc1 := indices{idx: idx, lidx: idx2 + len(w2)}
		ids = append(ids, oc1)

		lidx = oc1.lidx
	}


	var nd string = string(data)
	for _, s := range ids {
		nd = strings.ReplaceAll(string(nd), string(data[s.idx:s.lidx]), " ")
	}

	fmt.Println(getok(nd))
}

