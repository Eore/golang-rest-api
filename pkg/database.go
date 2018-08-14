package pkg

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost = "localhost"
	dbPort = "3306"
	dbUser = "root"
	dbName = "zweb"
)

func Connection() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", dbURL)
	return db
}

type Table struct {
	TableName string
	Fields    []string
	Props     []string
	Values    []interface{}
}

func (it *Table) ReadModel(model interface{}) {
	ref := reflect.TypeOf(model)
	tableName := ref.Name()
	fields := make([]string, ref.NumField())
	props := make([]string, ref.NumField())
	for i := 0; i < ref.NumField(); i++ {
		field, _ := ref.Field(i).Tag.Lookup("json")
		prop, _ := ref.Field(i).Tag.Lookup("sql")
		fields[i] = field
		props[i] = prop
	}
	it.TableName = tableName
	it.Fields = fields
	it.Props = props
}

func (it *Table) ReadValue(model interface{}) {
	ref := reflect.ValueOf((model))
	values := make([]interface{}, ref.NumField())
	for i := 0; i < ref.NumField(); i++ {
		values[i] = ref.Field(i).Interface()
	}
	it.Values = values
}

func (it *Table) Init(connection *sql.DB) {
	var str string
	var strs []string
	for i, val := range it.Fields {
		str = strings.Join([]string{val, it.Props[i]}, " ")
		strs = append(strs, str)
	}
	newStrs := "create table if not exists " + it.TableName + " (" + strings.Join(strs, ", ") + ")"
	Connection().Query(newStrs)
}
