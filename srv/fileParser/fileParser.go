package fileParser

import (
	"errors"
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
	Value      interface{}
}

// ParseFlatFile ...
func ParseFlatFile(formatFileName, dataFileName string) ([]Data, error) {
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

	// Check if the file is empty
	if len(bytes) == 0 {
		return res, errors.New("file is empty")
	}

	lines := strings.Split(string(bytes), "\n")

	var records []Data

	for _, line := range lines {
		if records, err = readDataLine(line, formats); err != nil {
			return res, err
		}

		res = append(res, records...)
	}

	return res, nil
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

	// Check if the file is empty
	if len(bytes) == 0 {
		return formats, errors.New("file is empty")
	}

	lines := strings.Split(string(bytes), "\n")

	var format Format

	for i := 1; i < len(lines); i++ {
		if format, err = readFormatLine(lines[i]); err != nil {
			return formats, err
		}

		formats = append(formats, format)
	}

	return formats, nil
}

// readFormatLine:
// sample: "name,10,TEXT"
func readFormatLine(line string) (Format, error) {
	var err error
	var format Format

	// check if line is empty
	if len(line) == 0 {
		return format, errors.New("empty string")
	}

	// remove all space
	line = strings.ReplaceAll(line, " ", "")

	words := strings.Split(line, ",")

	// check if there are only three columns
	if len(words) != 3 {
		return format, errors.New("invalid input format")
	}

	var intVar int

	if intVar, err = strconv.Atoi(words[1]); err != nil {
		return format, err
	}

	// check if the first and the third columns are empty
	if len(words[0]) == 0 || len(words[2]) == 0 {
		return format, errors.New("empty columns")
	}

	format = Format{
		ColumnName: words[0],
		Width:      intVar,
		DateType:   words[2],
	}

	return format, nil
}

// readDataLine:
// sample: "Diabetes  1  1"
func readDataLine(line string, formats []Format) ([]Data, error) {
	var res []Data
	var err error

	// check if line is empty
	if len(line) == 0 {
		return res, errors.New("empty line")
	}

	start := 0

	for _, format := range formats {
		if start+format.Width > len(line) {
			return res, errors.New("slice bounds out of range")
		}

		value := strings.ReplaceAll(line[start:start+format.Width], " ", "")

		if value == "" {
			return res, errors.New("empty value")
		}

		data := Data{
			ColumnName: format.ColumnName,
		}

		switch format.DateType {
		case "TEXT":
			data.Value = value
		case "BOOLEAN":
			if value == "1" {
				data.Value = true
			} else if value == "0" {
				data.Value = false
			} else {
				return res, errors.New("invalid boolean input value")
			}
		case "INTEGER":
			var intVar int
			if intVar, err = strconv.Atoi(value); err != nil {
				return res, err
			}
			data.Value = intVar
		default:
			return res, errors.New("type not exists")
		}

		res = append(res, data)

		start = start + format.Width
	}

	return res, nil
}
