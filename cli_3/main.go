package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: app <command> [options]")
		fmt.Println("Commands: greet, bye")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		name := greetCmd.String("name", "World", "Name to greet")
		greetCmd.Parse(os.Args[2:])
		fmt.Printf("Hello, %s!\n", *name)
	case "bye":
		byeCmd := flag.NewFlagSet("bye", flag.ExitOnError)
		name := byeCmd.String("name", "World", "Name to say bye to")
		byeCmd.Parse(os.Args[2:])
		fmt.Printf("Goodbye, %s!\n", *name)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
