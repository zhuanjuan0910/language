package main

import "fmt"

func main() {
	fn := func() string {
		return "hello world"
	}
	fmt.Println(fn())
	fc := []func(x int) int{
		func(x int) int { return x + 10 },
		func(x int) int { return x - 10 },
	}
	fmt.Println(fc[0](100))
	fmt.Println(fc[1](100))
	d := struct {
		fn func() string
	}{fn: func() string { return "how are you" }}
	fmt.Println(d.fn())
	fns := make(chan func() int, 2)
	fns <- func() int { return 10 }
	fmt.Println((<-fns)())
}
