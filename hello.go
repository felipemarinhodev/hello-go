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
	fmt.Scanf("%d", &command)

	fmt.Println("O comando escolhido foi", command)
}
