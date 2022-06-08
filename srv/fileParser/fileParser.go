package fileParser

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Format ...
type Format struct {
	ColumnName string
	Width      int
	DateType   string
}

// Data ...
type Data struct {
	ColumnName string
	Type       string
	Value      interface{}
}

// parseFormatFile ...
// columnName, width, datatype
func ParseFormatFile(filename string) ([]Format, error) {
	var bytes []byte
	var formats []Format
	var err error

	if bytes, err = ioutil.ReadFile(filename); err != nil {
		return formats, err
	}

	lines := strings.Split(string(bytes), "\n")

	for i := 1; i < len(lines); i++ {
		words := strings.Split(lines[i], ",")

		var intVar int
		if intVar, err = strconv.Atoi(words[1]); err != nil {
			return formats, err
		}

		formats = append(formats, Format{
			ColumnName: words[0],
			Width:      intVar,
			DateType:   words[2],
		})
	}

	return formats, nil
}

// ParseFlatfile ...
func ParseFlatfile(formatFileName, dataFileName string) ([]Data, error) {
	var err error
	var formats []Format
	var res []Data

	if formats, err = ParseFormatFile(formatFileName); err != nil {
		return res, err
	}

	var bytes []byte

	if bytes, err = ioutil.ReadFile(dataFileName); err != nil {
		return res, err
	}

	lines := strings.Split(string(bytes), "\n")

	for _, line := range lines {
		start := 0
		for _, format := range formats {
			res = append(res, Data{
				ColumnName: format.ColumnName,
				Type:       format.DateType,
				Value:      strings.ReplaceAll(line[start:start+format.Width], " ", ""),
			})

			start = start + format.Width
		}
	}

	return res, nil
}
