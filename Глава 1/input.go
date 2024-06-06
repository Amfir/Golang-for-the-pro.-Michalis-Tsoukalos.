package main

import "fmt"

func main() {
	// Получить пользовательский ввод
	fmt.Printf("please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}