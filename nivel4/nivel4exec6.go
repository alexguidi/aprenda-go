package main

import "fmt"

func main() {
	slice := make([]string, 3, 3)

	slice = []string{"Sao Paulo", "Minas", "Rio de Janeiro"}

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
