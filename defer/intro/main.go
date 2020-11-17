package main

import "fmt"

func DemoDefer() {
	
	f := 2

	defer fmt.Println("Execution Ended", f)

	defer fmt.Println("Second last statement")

	f += 2

	fmt.Println("Execution Started", f)
}

func main() {
	DemoDefer()

}