package main

import (
	"CloverHealth/srv/fileParser"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	var err error
	var data []fileParser.Data

	formatFileName := "../specs/fileformat1.csv"
	dataFileName := "../data/fileformat1_2022-06-07.txt"

	if data, err = fileParser.ParseFlatfile(formatFileName, dataFileName); err != nil {
		fmt.Printf("ERROR: %+v\r\n", errors.Wrap(err, ""))

		return
	}

	fmt.Println(data)

	for _, v := range data {
		switch v.Type {
		case "TEXT":
			fmt.Printf("Value: %#v\n", v.Value.(string))
		case "BOOLEAN":
			if v.Value == "1" {
				fmt.Printf("Value: %#v\n", true)
			} else if v.Value == "0" {
				fmt.Printf("Value: %#v\n", false)
			}
		case "INTEGER":
			intVar, _ := strconv.Atoi(v.Value.(string))
			fmt.Printf("Value: %#v\n", intVar)
		}
	}
}
