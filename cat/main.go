package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func PrintResult(str string) {
	for _, val := range str {
		z01.PrintRune(val)
	}
}
func MyReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "error"
	}
	return string(content)
}
func main() {
	args := os.Args[1:]
	finish := false
	for _, fileName := range args {
		if _, err := os.Stat(fileName); err != nil {
			PrintResult("ERROR: open " + fileName + ": no such file or directory\n")
			os.Exit(1)
			return
		}
		PrintResult(MyReadFile(fileName))
		finish = true
	}
	if !finish {
		reader := io.TeeReader(os.Stdin, os.Stdout)
		io.ReadAll(reader)
		os.Stdin.Close()
		os.Stdout.Close()
	}
}
