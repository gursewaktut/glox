package main

import (
	"bufio"
	"fmt"
	//"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64);
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt();
	}

}

func runFile(path string) {

	dat, err := os.ReadFile(path)
	check(err)
	run(string(dat))

}

func runPrompt() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		check(err)
		run(line)
	}
}

func run(source string) {
	fmt.Print(source)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
