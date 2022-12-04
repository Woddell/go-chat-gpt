package main

import (
	"flag"
	"fmt"
	"strings"
)

func init() {
	SetEnv()
}

func main() {
	message := getCliMessage()
	response := Chat(message)
	fmt.Println(response.Choices[0].Text)
}

func getCliMessage() string {
	_ = flag.String("", "", "pass any value and it'll be converted to a string")
	flag.Parse()
	args := flag.Args()
	return strings.Join(args, " ")
}
