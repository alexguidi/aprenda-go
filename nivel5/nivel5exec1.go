package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{nome: "Alex", sobrenome: "Guidi", sabores: []string{"morango", "chocolate"}}
	pessoa2 := pessoa{nome: "Nyc", sobrenome: "Fasolo", sabores: []string{"morango", "pistachio"}}

	fmt.Println(pessoa1)

	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}
}
