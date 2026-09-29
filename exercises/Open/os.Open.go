package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func task1() {
	_, err := os.Open("nope.csv")
	if err != nil {
		fmt.Printf("file is not found %v\n", err)
	}
	fmt.Println(errors.Is(err, fs.ErrNotExist))
	var pe *fs.PathError
	if errors.As(err, &pe) {
		fmt.Println(pe.Op, pe.Path)
	}

}
func main() {
	task1()
}
