package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: code-translator <source-file> <target-lang>")
		fmt.Println("Supported: python, javascript, rust")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()
	lang := os.Args[2]
	scanner := bufio.NewScanner(f)
	translated := 0
	fmt.Printf("Code Translator: -> %s\n\n", strings.Title(lang))
	for scanner.Scan() {
		line := scanner.Text()
		switch lang {
		case "python":
			line = strings.ReplaceAll(line, "func ", "def ")
			line = strings.ReplaceAll(line, "fmt.Println", "print")
		case "javascript":
			line = strings.ReplaceAll(line, "func ", "function ")
			line = strings.ReplaceAll(line, "fmt.Println", "console.log")
		case "rust":
			line = strings.ReplaceAll(line, "func ", "fn ")
			line = strings.ReplaceAll(line, "fmt.Println", "println!")
		}
		fmt.Println(line)
		translated++
	}
	fmt.Printf("\nTranslated %d lines to %s\n", translated, lang)
}