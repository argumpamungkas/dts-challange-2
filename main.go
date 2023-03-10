package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CHALLANGE 2

func main() {

	i := 0
	for i <= 4 {
		fmt.Printf("Nilai i = %d \n", i)
		i++
	}

	fmt.Println(strings.Repeat("=", 20))

	arrayUnicode := []string{"U+0421", "U+0410", "U+0428", "U+0410", "U+0420", "U+0412", "U+041E"}

	j := 0
	for j <= 10 {
		if j == 5 {
			for i, v := range arrayUnicode {
				position := i * 2
				value, _ := strconv.ParseInt(v[2:], 16, 16)
				fmt.Printf("character %#U starts at byte position %d \n", value, position)
			}
		} else {
			fmt.Printf("Nilai j = %d \n", j)
		}
		j++
	}
}
