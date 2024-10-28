package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	after := flag.Int("A", 0, "Print +N lines after match")
	before := flag.Int("B", 0, "Print +N lines until match")
	context := flag.Int("C", 0, "Print ±N lines around match")
	count := flag.Bool("c", false, "Count of lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Exclude matches")
	fixed := flag.Bool("F", false, "Exact string match, not a pattern")
	lineNum := flag.Bool("n", false, "Print line number")
	flag.Parse()

	var input []string

	if flag.NArg() == 0 {
		log.Fatal("Search pattern not specified")
	}

	pattern := flag.Arg(0)

	if flag.NArg() > 1 {
		filename := flag.Arg(1)
		lines, err := readLines(filename)
		if err != nil {
			log.Fatalf("error of reading file: %v\n", err)
		}
		input = lines
	} else {
		lines, err := readStdin()
		if err != nil {
			log.Fatalf("error of reading stdin: %v\n", err)
		}
		input = lines
	}

	if *context > 0 {
		*before = *context
		*after = *context
	}

	var matches []int

	for idx, line := range input {
		lineToCheck := line
		patternToCheck := pattern

		if *ignoreCase {
			lineToCheck = strings.ToLower(line)
			patternToCheck = strings.ToLower(pattern)
		}

		var isMatch bool
		if *fixed {
			isMatch = lineToCheck == patternToCheck
		} else {
			matched, err := regexp.MatchString(patternToCheck, lineToCheck)
			if err != nil {
				log.Printf("error of matching: %v\n", err)
			}
			isMatch = matched
		}

		if *invert {
			isMatch = !isMatch
		}

		if isMatch {
			matches = append(matches, idx)
		}
	}

	if *count {
		fmt.Println(len(matches))
		return
	}

	linesToPrint := make(map[int]bool)

	for _, idx := range matches {
		start := idx - *before
		if start < 0 {
			start = 0
		}
		end := idx + *after
		if end >= len(input) {
			end = len(input) - 1
		}
		for i := start; i <= end; i++ {
			linesToPrint[i] = true
		}
	}

	for i, line := range input {
		if linesToPrint[i] {
			if *lineNum {
				fmt.Printf("%d:%s\n", i+1, line)
			} else {
				fmt.Println(line)
			}
		}
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readStdin() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
