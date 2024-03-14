package main

//editional push
import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

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

	if !(*count || *unique || *duplicate) {
		printUniqueLines(lines, *numFields, *numChars, *ignoreCase)
	} else if *count {
		removeDuplicates(lines, "c", *numFields, *numChars, *ignoreCase)
	} else if *duplicate {
		removeDuplicates(lines, "d", *numFields, *numChars, *ignoreCase)
	} else if *unique {
		removeDuplicates(lines, "u", *numFields, *numChars, *ignoreCase)
	} //тут наверное не получится switch
}

func printUniqueLines(lines []string, nf int, nC int, iC bool) {
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
		fmt.Println(line)
	}
}
func removeDuplicates(lines []string, mode string, nf int, nC int, iC bool) {
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
			printStringscount(c, lines[i-1], mode)
			c = 1
		}
	}
	printStringscount(c, lines[len(lines)-1], mode)
}
func printStringscount(c int, str1 string, mode string) {
	if mode == "c" {
		fmt.Println(c, str1)
	} else if mode == "d" {
		if c > 1 {
			fmt.Println(str1)
		}
	} else {
		if c == 1 {
			fmt.Println(str1)
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
func newFunc() {
	fmt.Println("uu")
}
