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
	dbPass := os.Getenv("DB_PASS")
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", dbURL)
	return db
}

type Table struct {
	tableName string
	fields    []string
	props     []string
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
	it.tableName = tableName
	it.fields = fields
	it.props = props
}

func (it *Table) Init(connection *sql.DB) {
	var str string
	var strs []string
	for i, val := range it.fields {
		str = strings.Join([]string{val, it.props[i]}, " ")
		strs = append(strs, str)
	}
	newStrs := "create table if not exists " + it.tableName + " (" + strings.Join(strs, ", ") + ")"
	Connection().Query(newStrs)
}
