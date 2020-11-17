package main

func FuncWithoutDefer() {
	sum := 0

	for i := 1; i < 100; i++ {
		sum += i
	}

	sum += 10

}

func FuncWithDefer() {

	sum := 0

	defer func() {
		sum += 10
	}()

	for i := 1; i < 100; i++ {
		sum += i
	}

}
 

func main() {

}