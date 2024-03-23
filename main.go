package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	inputFilePath := flag.String("input", "", "Файл ввода (оставьте пустым для stdin)")
	outputFilePath := flag.String("output", "", "Файл вывода (оставьте пустым для stdout)")

	count := flag.Bool("c", false, "подсчитать количество встречаний строки")
	duplicate := flag.Bool("d", false, "вывести повторяющиеся строки")
	unique := flag.Bool("u", false, "вывести неповторяющиеся строки")
	ignoreCase := flag.Bool("i", false, "не учитывать регистр букв")
	numFields := flag.Int("f", 0, "не учитывать первые num_fields полей")
	numChars := flag.Int("s", 0, "не учитывать первые num_chars символов")

	flag.Parse()

	if (*count && *duplicate) || (*count && *unique) || (*duplicate && *unique) {
		fmt.Println("Флаги -c, -d, и -u не могут быть использованы вместе. ")
		return
	}

	var reader io.Reader = os.Stdin
	var writer io.Writer = os.Stdout

	if *inputFilePath != "" {
		file, err := os.Open(*inputFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при открытии файла ввода: %v\n", err)
			return
		}
		defer file.Close()
		reader = file
	}

	if *outputFilePath != "" {
		file, err := os.Create(*outputFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при создании файла вывода: %v\n", err)
			return
		}
		defer file.Close()
		writer = file
	}

	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	switch {
	case *count:
		removeDuplicates(writer, lines, "c", *numFields, *numChars, *ignoreCase)
	case *duplicate:
		removeDuplicates(writer, lines, "d", *numFields, *numChars, *ignoreCase)
	case *unique:
		removeDuplicates(writer, lines, "u", *numFields, *numChars, *ignoreCase)
	default:
		printUniqueLines(writer, lines, *numFields, *numChars, *ignoreCase)
	}

}

func printUniqueLines(writer io.Writer, lines []string, nf int, nC int, iC bool) {
	result := []string{lines[0]}
	for i := 1; i < len(lines); i++ {
		s1 := ignoreNFields(lines[i], nf)
		s2 := ignoreNFields(lines[i-1], nf)
		if nC > len(s1) || nC > len(s2) {
			nC = 0
		}
		if iC {
			s1 = strings.ToLower(s1)
			s2 = strings.ToLower(s2)
		}
		if s1[nC:] != s2[nC:] {
			result = append(result, lines[i])
		}
	}
	for _, line := range result {
		fmt.Fprintln(writer, line)
	}
}
func removeDuplicates(writer io.Writer, lines []string, mode string, nf int, nC int, iC bool) {
	c := 1
	for i := 1; i < len(lines); i++ {
		s1 := ignoreNFields(lines[i], nf)
		s2 := ignoreNFields(lines[i-1], nf)
		if nC > len(s1) || nC > len(s2) {
			nC = 0
		}
		if iC {
			s1 = strings.ToLower(s1)
			s2 = strings.ToLower(s2)
		}
		if s1[nC:] == s2[nC:] {
			c += 1
		} else {
			printStringscount(writer, c, lines[i-1], mode)
			c = 1
		}
	}
	printStringscount(writer, c, lines[len(lines)-1], mode)
}
func printStringscount(writer io.Writer, c int, str1 string, mode string) {
	if mode == "c" {
		fmt.Fprintln(writer, c, str1)
	} else if mode == "d" {
		if c > 1 {
			fmt.Fprintln(writer, str1)
		}
	} else {
		if c == 1 {
			fmt.Fprintln(writer, str1)
		}
	}
}
func ignoreNFields(s string, n int) string {
	fields := strings.Fields(s)
	if n >= len(fields) {
		return ""
	}
	return strings.Join(fields[n:], " ")
}
