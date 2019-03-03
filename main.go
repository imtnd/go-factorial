package main

import (
	"flag"
	"fmt"
	"github.com/imtnd/go-factorial/pkg/factorial"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.NArg()
	if args != 1 {
		fmt.Printf("args is %#v.\n", args)
		fmt.Println("args must be 1.")
		return
	}
	arg1, _ := strconv.Atoi(flag.Arg(0))

	result := factorial.GetFactorial(arg1)
	fmt.Println("Factorial is", result)
}
