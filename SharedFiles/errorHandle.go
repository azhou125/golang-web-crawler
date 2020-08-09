package SharedFiles

import "fmt"

func CheckError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
	}
}