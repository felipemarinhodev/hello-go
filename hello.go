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

	switch command {
		case 1:
			fmt.Println("Monitorando...")
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
		default:
			fmt.Println("Comando não suportado.")
	}
}
