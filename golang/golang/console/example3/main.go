package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string") //括号里分别是命令行里的可选参数，默认值，描述
	boolPtr := flag.Bool("fork", false, "a bool")
	numPtr := flag.Int("num", 43, "an int")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	//所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析。
	flag.Parse()
	/*
		这里我们将仅输出解析的选项以及后面的位置参数。
		注意，我们需要使用类似 *wordPtr 这样的语法来对指针解引用，从而得到选项的实际值。
	*/
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
