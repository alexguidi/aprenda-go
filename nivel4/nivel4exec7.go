package main

import "fmt"

func main() {
	ss := [][]string{
		[]string{"Alex", "Guidi", "Programming"},
		[]string{"Nyc", "Fasolo", "Dance"},
		[]string{"Bruno", "Guidi", "Music"},
	}

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}
}
