package main

import (
	"fmt"
	"reflect"
)

func main()  {
	name := "Felipe"
	age := 24
	version := 1.1
	fmt.Println("olá sr.", name, "sua idade é", age)
	fmt.Println("este programa está na versão", version)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(name), "e idade é", reflect.TypeOf(age), "e a versão", reflect.TypeOf(version))
}
