package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	lineRowKey int = iota
	lineFamily
	lineValue
)

type Table struct {
	Meta Meta
	Rows []Row
}

type Meta struct {
	Family []string
	Column map[string]string // map[familyName]columnName
}

type Row struct {
	Key   string
	Value map[string]string // map[family:identify@timestamp]value
}

func checkLine(text string) int {
	if strings.HasPrefix(text, "    ") {
		return lineValue
	} else if strings.HasPrefix(text, "  ") {
		return lineFamily
	} else {
		return lineRowKey
	}
}

func parseRowKey(text string) Row {
	return Row{
		Key:   strings.TrimSpace(text),
		Value: make(map[string]string, 0),
	}
}

// TODO Add meta data
func parseColumnKey(text string) string {
	text = strings.TrimSpace(text)
	return text
	/*
		splitIndex := strings.LastIndex(text, "@")
		if splitIndex == -1 {
			panic(fmt.Errorf("not found @"))
		}
		familyColumn := text[:splitIndex]
		time := time.ParseDuration(text[splitIndex+1:])
		family, column := strings.Split(familyColumn, ":")
	*/
}

func parseValue(text string) string {
	return strings.TrimSpace(text)
}

func Parse(r io.Reader) (*Table, error) {
	line := strings.Repeat("-", 40)
	table := Table{}
	currentColumnKey := ""

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == line {
			continue
		}
		switch checkLine(text) {
		case lineRowKey:
			table.Rows = append(table.Rows, parseRowKey(text))
		case lineFamily:
			currentColumnKey = parseColumnKey(text)
			table.Rows[len(table.Rows)-1].Value[currentColumnKey] = ""
		case lineValue:
			table.Rows[len(table.Rows)-1].Value[currentColumnKey] = parseValue(text)
		default:
			panic(fmt.Errorf("no check pattern"))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error occured to scan: %v", err)
	}
	return &table, nil
}
