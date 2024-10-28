package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "", "Select fields (columns)")
	delimiter := flag.String("d", "\t", "Use a different delimiter")
	separated := flag.Bool("s", false, "Only delimited lines")
	flag.Parse()

	if *fields == "" {
		fmt.Println("Option requires an argument -- f")
		flag.Usage()
		os.Exit(1)
	}

	fieldIndices, err := parseFields(*fields)
	if err != nil {
		log.Fatalf("error of parsing fields: %v\n", err)
	}

	lines, err := readStdin()
	if err != nil {
		log.Fatalf("error of reading stdin: %v\n", err)
	}

	processLines(lines, fieldIndices, *delimiter, *separated)
}

func parseFields(fieldsStr string) ([]int, error) {
	fieldsParts := strings.Split(fieldsStr, ",")
	var fields []int
	for _, part := range fieldsParts {
		fieldNum, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil || fieldNum < 1 {
			return nil, fmt.Errorf("incorrect field value: %s", part)
		}
		fields = append(fields, fieldNum-1)
	}
	return fields, nil
}

func readStdin() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func processLines(lines []string, fields []int, delimiter string, separated bool) {
	for _, line := range lines {
		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		columns := strings.Split(line, delimiter)

		var selectedColumns []string
		for _, fieldIdx := range fields {
			if fieldIdx >= 0 && fieldIdx < len(columns) {
				selectedColumns = append(selectedColumns, columns[fieldIdx])
			}
		}
		fmt.Println(strings.Join(selectedColumns, delimiter))
	}
}
