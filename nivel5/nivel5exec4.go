package main

import "fmt"

func main() {

	test := struct {
		meuMap   map[string]string
		meuSlice []string
	}{
		map[string]string{"Alex": "test"},
		[]string{"amarelo", "azul"},
	}

	fmt.Println(test)
}
