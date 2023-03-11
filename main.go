package main

import (
	"fmt"
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

	charUnicode := "САШАРВО"

	j := 0
	for j <= 10 {
		if j == 5 {
			for i, v := range charUnicode {
				fmt.Printf("character %#U starts at byte position %d\n", v, i)
			}
		} else {
			fmt.Printf("Nilai j = %d \n", j)
		}
		j++
	}
}
