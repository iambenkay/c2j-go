package c2j_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	parser, _ := Load("./samples/data.csv", true)

	assert.Equal(t, []string{"number", "spelling"}, parser.header, "They should be equal")
	assert.Equal(t, [][]string{{"1", "one"}, {"2", "two"}, {"3", "three"}}, parser.rows, "They should be equal")
}

func TestLoadWithoutHeaders(t *testing.T) {
	parser, _ := Load("./samples/data_without_headers.csv", false)

	assert.Equal(t, 0, len(parser.header), "Header should be empty")
	assert.Equal(t, [][]string{{"1", "one"}, {"2", "two"}, {"3", "three"}}, parser.rows, "They should be equal")
}

func TestParser_ToMap(t *testing.T) {
	parser, _ := Load("./samples/data.csv", true)

	result := parser.ToMap()

	assert.Equal(t, []map[string]interface{}{
		{"number": int64(1), "spelling": "one"},
		{"number": int64(2), "spelling": "two"},
		{"number": int64(3), "spelling": "three"},
	}, result, "They should be equal")
}

func TestParser_ToMapWithoutHeaders(t *testing.T) {
	parser, _ := Load("./samples/data_without_headers.csv", false)

	result := parser.ToMap()

	assert.Equal(t, []map[string]interface{}{
		{"col1": int64(1), "col2": "one"},
		{"col1": int64(2), "col2": "two"},
		{"col1": int64(3), "col2": "three"},
	}, result, "They should be equal")
}

func TestParser_ToJSON(t *testing.T) {
	parser, _ := Load("./samples/data.csv", true)

	result, _ := parser.ToJSON()

	assert.JSONEq(t, `[
		{"number": 1, "spelling": "one"},
		{"number": 2, "spelling": "two"},
		{"number": 3, "spelling": "three"}
	]`, result, "They should be equal")
}

func TestParser_ToJSONWithoutHeaders(t *testing.T) {
	parser, _ := Load("./samples/data_without_headers.csv", false)

	result, _ := parser.ToJSON()

	assert.JSONEq(t, `[
		{"col1": 1, "col2": "one"},
		{"col1": 2, "col2": "two"},
		{"col1": 3, "col2": "three"}
	]`, result, "They should be equal")
}
