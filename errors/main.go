package main

import (
	"errors"
	"fmt"
)

type errComment error

func main() {
	err := commentError()

	switch err.(type) {
	case errComment:
		fmt.Println("Error Comment")
	default:
		fmt.Println("Normal Error")
	}

}

func commentError() error {
	return errors.New("comment error")
}
