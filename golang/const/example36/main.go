package main

import "strconv"

type xyz struct {
	age int
}

func (ab xyz) test(x string) string {
	return strconv.Itoa(ab.age) + x
}

func main() {
	a := xyz{12}
	a.test("casdc")
}
