package main

import "fmt"

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func main() {

	cliente := Cliente{
		Nome:  "Tiago",
		Email: "tfc@gmail.com",
		CPF:   10515156651,
	}

	fmt.Println(cliente)

}

// struct no go é como herança em outras linguages
// tem uma struct que tem um metodo 