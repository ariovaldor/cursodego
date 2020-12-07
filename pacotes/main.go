package main

import (
	"fmt"
	"pacotes/operadora"
	"pacotes/prefixo"
	"strconv"
)

var NomedoUsuario = "Ariovaldo Full Stack"
var identOperadora string

func main() {
	fmt.Printf("Nome do Usuário: %s\r\n", NomedoUsuario)
	fmt.Printf("Prefixo: %d\r\n", prefixo.Capital)
	identOperadora = "(" + strconv.Itoa(prefixo.Capital) + ") " + operadora.Operadora
	fmt.Printf("Operadora: %s\r\n", identOperadora)

}
