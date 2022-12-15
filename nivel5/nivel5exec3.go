package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	caminhonete := caminhonete{tracaoNasQuatro: true, veiculo: veiculo{portas: 2, cor: "preta"}}

	sedan := sedan{
		modeloLuxo: true,
		veiculo: veiculo{
			portas: 4,
			cor:    "amarela",
		},
	}

	fmt.Println(caminhonete)
	fmt.Println(sedan)
	fmt.Println(caminhonete.portas)
	fmt.Println(caminhonete.cor)
}
