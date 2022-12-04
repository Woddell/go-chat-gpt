package main

import "os"

func LogToFile(content string) {
	f, err := os.OpenFile("main.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}
}
