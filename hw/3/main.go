package main

import (
	"calculations/calculations"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//var n int64 = 3
	// itog := calculations.Calculate(n, true)
	// fmt.Printf("%d! = %d\n", n, itog)
	log := flag.Bool("log", false, "Do loging or not")
	flag.Parse()
	var args []string
	if *log {
		args = os.Args[2:]
	} else {
		args = os.Args[1:]
	}
	for _, a := range args {
		int1, err := strconv.ParseInt(a, 10, 64)
		if err == nil {
			fmt.Printf("%s! = %d\n", a, calculations.Calculate(int1, *log))
		} else {
			fmt.Printf("Incorrect input: %s\n", a)
			os.Exit(1)
		}

	}

}
