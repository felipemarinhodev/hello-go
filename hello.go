package main

import (
	"fmt"
)

func main()  {
	name := "Felipe"
	version := 1.1
	fmt.Println("olá sr.", name)
	fmt.Println("este programa está na versão", version)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa.")

	var command int;
	fmt.Scan(&command)

	fmt.Println("O comando escolhido foi", command)

	if command == 1 {
		fmt.Println("Monitorando...")
	} else if command == 2 {
		fmt.Println("Exibindo Logs...")
	} else if command == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Comando não suportado.")
	}
}
