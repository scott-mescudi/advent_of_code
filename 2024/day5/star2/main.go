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

func parseData(filename string) ([][]string, []mcnugget) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	m1 := []mcnugget{}
	m2 := [][]string{}

	scanner := bufio.NewScanner(file)
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

	return m2, m1
}

func filter(nums []string, combos []mcnugget) {
    for {
        swapped := false
        nmap := map[string]int{}
        for index, x := range nums {
            nmap[x] = index
        }

        for i, v := range nums {
            for _, combo := range combos {
                if combo.a == v {
                    if idx, ok := nmap[combo.b]; ok && idx < i {
                        nums[i], nums[idx] = nums[idx], nums[i]
                        swapped = true
                        break
                    }
                }
            }
        }

        if !swapped || check(nums, combos) {
            break
        }
    }
}


func main() {
	m2, m1 := parseData("../data.txt")
	ss := [][]string{}
	for _, row := range m2[1:] {
		if ok := check(row, m1); !ok {
			ss = append(ss, row)
		}
	}

	var total int = 0

	for _, rw := range ss {
		filter(rw, m1)
	}

	for _, r := range ss {
		n, err := strconv.Atoi(r[(len(r)/2)])
		if err != nil {
			log.Fatalln(err)
		}
		total+=n
	}

	fmt.Println(total)
}