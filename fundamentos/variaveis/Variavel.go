package main

import "fmt"

var (
	//Endereco é um valor importante e tem de ser publico
	Endereco string //endereco = ""
	//Telefone é um valor importante para a nossa aula
	Telefone            string = "+55081999999999"
	quantidade, estoque int     //quantidade = 0
	comprou             bool    //comprou = false
	valor               float64 //valor = 0.00
	palavras            rune
)

func main() {

	teste := "Valor de teste"
	fmt.Printf("Endereço: %s\r\n", Endereco)
	fmt.Printf("Telefone: %s\r\n", Telefone)
	fmt.Printf("Quantidade: %d\r\n", quantidade)
	fmt.Printf("Comprou: %v\r\n", comprou)
	fmt.Printf("Palavras: %v\r\n", palavras)
	fmt.Printf("O valor de teste é: %v\r\n", teste)

}
