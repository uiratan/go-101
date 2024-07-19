package main

import "fmt"

var (
	name string
	N1   int
	n2   int32
)

func main() {
	mensagem := "Class #3 - Go 101"
	fmt.Println(mensagem)

	var a, b, c = true, 2.3, "hello"
	fmt.Println(a, b, c)

	var total int
	total++
	fmt.Println("total: ", total)

	name = "Uiratan Cavalcante"
	fmt.Println("Hello", name, "World!")

	var x, y = 10, 20
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

}
