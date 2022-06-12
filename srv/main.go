package main

import (
	"CloverHealth/srv/fileParser"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	var err error
	var data []fileParser.Data

	formatFileName := "../specs/fileformat1.csv"
	dataFileName := "../data/fileformat1_2022-06-07.txt"

	if data, err = fileParser.ParseFlatFile(formatFileName, dataFileName); err != nil {
		fmt.Printf("ERROR: %+v\r\n", errors.Wrap(err, ""))

		return
	}

	//fmt.Printf("%#v\n", data)
	fmt.Println(data)

	for _, v := range data {
		switch v.Value.(type) {
		case string:
			fmt.Printf("Type: %T, val: %v\n", v.Value, v.Value)
		case bool:
			fmt.Printf("Type: %T, val: %v\n", v.Value, v.Value)
		case int:
			fmt.Printf("Type: %T, val: %v\n", v.Value, v.Value)
		}
	}
}
