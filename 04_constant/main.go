package main

import "fmt"

const age = 34

var name2 = "sahil"

func main() {
	const name = "golang" //can not be reassigned

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(name2)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port,host)
}