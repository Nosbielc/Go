package main

import (
	"Go/Go/fundamentos/prefixo"
	"fmt"
)

import "operadora/nome.go"

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)
}
