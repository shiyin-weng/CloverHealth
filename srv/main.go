package main

import (
	"CloverHealth/srv/fileParser"
	"fmt"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	var data []fileParser.Data
	var formats []fileParser.Format

	//formatFileName := "../specs/fileformat1.csv"
	//dataFileName := "../data/fileformat1_2022-06-07.txt"
	formatFileName := "../specs/fileformat2.csv"
	dataFileName := "../data/fileformat2_2022-06-12.txt"

	if formats, data, err = fileParser.ParseFlatFile(formatFileName, dataFileName); err != nil {
		fmt.Printf("ERROR: %+v\r\n", errors.Wrap(err, ""))

		return
	}

	//fmt.Printf("%#v\n", data)
	fmt.Println(formats)
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

	// DB operations
	//if err = mysqldb.InitDB(); err != nil {
	//	fmt.Printf("ERROR: %+v\r\n", errors.Wrap(err, ""))

	//	return
	//}

	//if err = mysqldb.CreateTable(formatFileName, formats); err != nil {
	//	fmt.Printf("ERROR: %+v\r\n", errors.Wrap(err, ""))

	//	return
	//}

	//fmt.Println(data)

	//lines := strings.Split(formatFileName, "/")
	//line := lines[len(lines)-1]
	//tableName := line[:len(line)-4]

	//for i := 0; i < len(data); i += len(formats) {
	//	if err = mysqldb.InsertData(tableName, formats, data[i:i+len(formats)]); err != nil {
	//		fmt.Printf("ERROR: %+v\r\n", errors.Wrap(err, ""))

	//		return
	//	}
	//}
}
