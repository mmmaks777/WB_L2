package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	column  = flag.Int("k", 0, "Specifying the column to be sorted (words in a line can act as columns; the default separator is space)")
	numeric = flag.Bool("n", false, "Sort by numeric value.")
	reverse = flag.Bool("r", false, "Sort in reverse order.")
	unique  = flag.Bool("u", false, "Unique keys.")
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-sort [options] inputfile outputfile")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	lines, err := readLines(inputFile)
	if err != nil {
		log.Fatalf("error of reading file: %v\n", err)
	}

	if *unique {
		lines = removeDublicates(lines)
	}

	sortLines(lines)

	err = writeLines(outputFile, lines)
	if err != nil {
		log.Fatalf("error of writing lines in output file: %v\n", err)
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

func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func removeDublicates(lines []string) []string {
	seen := make(map[string]bool)
	uniqueLines := []string{}
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}

func sortLines(lines []string) {
	sort.SliceStable(lines, func(i, j int) bool {
		var less bool
		if *column > 0 {
			less = compareByColumn(lines[i], lines[j], *column-1)
		} else {
			less = lines[i] < lines[j]
		}
		if *reverse {
			return !less
		}
		return less
	})
}

func compareByColumn(a, b string, col int) bool {
	partsA := strings.Fields(a)
	partsB := strings.Fields(b)

	if col < len(partsA) && col < len(partsB) {
		if *numeric {
			numA, errA := strconv.ParseFloat(partsA[col], 64)
			numB, errB := strconv.ParseFloat(partsB[col], 64)
			if errA == nil && errB == nil {
				return numA < numB
			}
		}
		return partsA[col] < partsB[col]
	}
	return a < b
}
