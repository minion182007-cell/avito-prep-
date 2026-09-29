package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

var (
	ErrEmpty      = errors.New("empty value")
	ErrOutOfRange = errors.New("age out of range")
)

type ParseError struct {
	Line  int
	Value string
	Err   error
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("Line %d Value %q Err %v", p.Line, p.Value, p.Err)

}
func (p *ParseError) Unwrap() error {
	return p.Err
}
func parseAges(vals []string) ([]int, error) {
	var result []int
	for i, v := range vals {

		n, err := parseAge(v)
		if err != nil {
			return nil, &ParseError{Line: i + 1, Value: v, Err: err}
		}
		result = append(result, n)
	}

	return result, nil
}
func parseAge(s string) (int, error) {

	if s == "" {
		return 0, ErrEmpty
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parseAge %q: %w", s, err)
	}
	if 0 > num || num > 150 {
		return 0, fmt.Errorf("parseAge %d: %w", num, ErrOutOfRange)
	}
	return num, nil
}
func task4() {
	_, err := parseAges([]string{"23", "31", "abc", "40"})
	fmt.Println(err)

	var pe *ParseError
	if errors.As(err, &pe) {
		fmt.Println("line:", pe.Line)
	}
	fmt.Println(errors.Is(err, strconv.ErrSyntax))
}
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

func readConfig(path string) error {
	p, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("readConfig %s, %w", path, err)

	}
	defer p.Close()
	return nil
}
func loadApp() error {
	err := readConfig("nope.csv")
	if err != nil {
		return fmt.Errorf("loadApp: %w", err)
	}
	return nil
}
func task2() {
	err := loadApp()
	fmt.Println(errors.Is(err, fs.ErrNotExist))
	fmt.Println(err)
}
func task3() {
	for _, s := range []string{"23", "", "abc", "200", "-5"} {
		n, err := parseAge(s)
		fmt.Println(n, err,
			errors.Is(err, ErrEmpty),
			errors.Is(err, ErrOutOfRange),
			errors.Is(err, strconv.ErrSyntax))
	}
}
func main() {
	task1()
	task2()
	task3()
	task4()
}
