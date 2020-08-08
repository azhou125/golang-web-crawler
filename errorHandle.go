package main

import "fmt"

func checkError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
	}
}