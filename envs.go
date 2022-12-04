package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SetEnv() {
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		os.Setenv(parts[0], parts[1])
	}
}
