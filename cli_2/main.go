package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	name := flag.String("name", "World", "Who to greet")
	loud := flag.Bool("loud", false, "Shout the greeting")
	times := flag.Int("times", 1, "How many times")

	flag.Parse()

	msg := fmt.Sprintf("Hello, %s!", *name)

	if *loud {
		msg = strings.ToUpper(msg)
	}
	for i := 0; i < *times; i++ {
		fmt.Println(msg)
	}
}
