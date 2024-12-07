package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type mcnugget struct {
	a, b string
}


func check(nums []string, combos []mcnugget) bool {
	fin := false

	nmap := map[string]int{}
	

	for index, x := range nums {
		nmap[x] = index
	}

	
	for i, v := range nums {
		for _, combo := range combos {
			if combo.a == v {
				if idx, ok := nmap[combo.b]; ok {
					if idx < i {
						return false
					}else{ fin = true }
				}
			}
		}
	}

	return fin
}

func main() {
	file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m1 := []mcnugget{}
	m2 := [][]string{}
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.Contains(txt, "|"){
			parts := strings.Split(txt, "|")

			m1 = append(m1, mcnugget{a: parts[0], b: parts[1]})
		}else{
			cont := []string{}
			parts := strings.Split(txt, ",")
			for _, v := range parts {
				cont = append(cont, v)
			}

			if len(cont) == 0{
				continue
			}

			m2 = append(m2, cont)
		}
	}


	ss := [][]string{}
	for _, row := range m2[1:] {
		if ok := check(row, m1); ok {
			ss = append(ss, row)
		}
	}

	total := 0
	for _, r := range ss {
		size := len(r)
		num := r[(size/2)]
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalln(err)
		}
		total+=n
	}

	fmt.Println(total)
}