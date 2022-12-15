package main

import "fmt"

//65 a 90

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%U '%c' \n", i, i)
		}
	}
}
