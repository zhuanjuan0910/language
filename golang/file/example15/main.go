package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		fmt.Println("this is first defer")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Panic("this is  log panic")
	defer func() {
		fmt.Println("this is second defer")
	}()
}
