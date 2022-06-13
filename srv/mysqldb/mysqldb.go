package mysqldb

import (
	"CloverHealth/srv/fileParser"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

var DB *sql.DB

// InitDB ...
func InitDB() error {
	var err error

	// connect to database
	if DB, err = sql.Open("mysql", "sweng:password@/CloverHealth"); err != nil {
		return err
	}

	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	return nil
}

// CreateTable ...
func CreateTable(formatFileName string, formats []fileParser.Format) error {
	lines := strings.Split(formatFileName, "/")
	line := lines[len(lines)-1]
	tableName := line[:len(line)-4]

	//fmt.Println(line)

	schema := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %v (
		id INT NOT NULL AUTO_INCREMENT,
	`,
		tableName,
	)

	for _, format := range formats {
		switch format.DateType {
		case "TEXT":
			schema += fmt.Sprintf(`%v VARCHAR(255) NOT NULL, `, format.ColumnName)
		case "BOOLEAN":
			schema += fmt.Sprintf(`%v TINYINT NOT NULL, `, format.ColumnName)
		case "INTEGER":
			schema += fmt.Sprintf(`%v INT NOT NULL, `, format.ColumnName)
		default:
			return errors.New("type not exists")
		}
	}

	//name VARCHAR(255) NOT NULL,
	//valid TINYINT NOT NULL,
	//count INT NOT NULL,

	schema += `
		PRIMARY KEY ( id )
	);`

	//fmt.Println(schema)

	if _, err := DB.Exec(schema); err != nil {
		return err
	}

	return nil
}

// InsertData ...
func InsertData(tableName string, formats []fileParser.Format, data []fileParser.Data) error {
	columns := []string{}

	for _, format := range formats {
		columns = append(columns, format.ColumnName)
	}

	arr := []string{}
	for _, v := range data {
		switch v.Value.(type) {
		case string:
			arr = append(arr, fmt.Sprintf(`'%v'`, v.Value))
		case bool:
			arr = append(arr, fmt.Sprintf(`%v`, v.Value))
		case int:
			arr = append(arr, fmt.Sprintf(`%v`, v.Value))
		default:
			return errors.New("type not exists")
		}
	}

	sqlQuery := fmt.Sprintf(`
	INSERT INTO %v (%v) VALUES (%v); `,
		tableName,
		strings.Join(columns, ", "),
		strings.Join(arr, ", "),
	)

	fmt.Println(sqlQuery)
	return nil

	if _, err := DB.Exec(sqlQuery); err != nil {
		return err
	}

	return nil
}
