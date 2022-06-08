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

//func main() {
//	var fields []reflect.StructField
//	fields = append(fields, reflect.StructField{
//		Name: "Height",
//		Tag:  `json:"height"`,
//	})
//
//	fields[0].Type = reflect.TypeOf(float64(0))
//
//	fields = append(fields, reflect.StructField{
//		Name: "Age",
//		Tag:  `json:"age"`,
//	})
//
//	//Type: reflect.TypeOf(int(0)),
//	fields[1].Type = reflect.TypeOf(int(0))
//
//	typ := reflect.StructOf(fields)
//
//	v := reflect.New(typ).Elem()
//	v.Field(0).SetFloat(0.4)
//	v.Field(1).SetInt(2)
//	s := v.Addr().Interface()
//
//	w := new(bytes.Buffer)
//	if err := json.NewEncoder(w).Encode(s); err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("value: %+v\n", s)
//	fmt.Printf("json:  %s", w.Bytes())
//
//	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))
//	if err := json.NewDecoder(r).Decode(s); err != nil {
//		panic(err)
//	}
//	fmt.Printf("value: %+v\n", s)
//}
