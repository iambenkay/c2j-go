package c2j_go

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	includesHeader bool
	header         []string
	rows           [][]string
}

func Load(path string, includesHeader bool) (parser Parser, err error) {
	data, err := ioutil.ReadFile(path)
	reader := csv.NewReader(strings.NewReader(string(data)))
	parser = Parser{includesHeader: includesHeader}

	records, err := reader.ReadAll()
	if err == nil {
		if includesHeader {
			parser.header = records[0]
			parser.rows = records[1:]
		} else {
			parser.rows = records
		}
	}
	return
}

func (parser *Parser) ToMap() (result []map[string]interface{}) {
	var columnCount int

	if parser.includesHeader {
		columnCount = len(parser.header)
	} else {
		columnCount = len(parser.rows[0])
	}

	for _, row := range parser.rows {
		item := make(map[string]interface{})

		for i := 0; i < columnCount; i++ {
			var key string
			if parser.includesHeader {
				key = parser.header[i]
			} else {
				key = fmt.Sprintf("col%d", i+1)
			}
			item[key] = process(row[i])
		}

		result = append(result, item)
	}

	return
}

func (parser *Parser) ToJSON() (result string, err error) {
	data := parser.ToMap()

	resultBytes, err := json.Marshal(data)

	return string(resultBytes), err
}

func process(str string) (parsed interface{}) {
	var err error
	if match, _ := regexp.MatchString("^[1-9]\\d*\\.\\d+$", str); match {
		parsed, err = strconv.ParseFloat(str, 32)
		if err == nil {
			return
		}
	} else if match, _ := regexp.MatchString("^[1-9]\\d*$", str); match {
		parsed, err = strconv.ParseInt(str, 10, 32)
		if err == nil {
			return parsed
		}
	}
	parsed, err = strconv.ParseBool(str)
	if err == nil {
		return
	}
	return str
}
