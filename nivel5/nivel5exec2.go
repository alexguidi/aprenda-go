package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoas := make(map[string]pessoa)

	pessoa1 := pessoa{nome: "Alex", sobrenome: "Guidi", sabores: []string{"morango", "chocolate"}}
	pessoa2 := pessoa{nome: "Nyc", sobrenome: "Fasolo", sabores: []string{"morango", "pistachio"}}

	pessoas["Guidi"] = pessoa1
	pessoas["Fasolo"] = pessoa2

	for _, v := range pessoas {
		fmt.Println("Nome: ", v.nome)
		for _, v := range v.sabores {
			fmt.Println(v)
		}
	}

}
