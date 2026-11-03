package main

import (
	"errors"
	"fmt"
)

var ErrNil = errors.New("is nil")

type NilErr struct{}

func (e *NilErr) Error() string {
	return "is nil"
}

func RetutNil() error {

	return nil
}
func main() {
	err := RetutNil()
	fmt.Println(err == nil)
}
