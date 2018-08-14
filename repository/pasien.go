package repository

import (
	"fmt"

	"../model"
	"../pkg"
)

var t pkg.Table

func init() {
	fmt.Println("init table")
	t.ReadModel(model.Pasien{})
	t.Init(pkg.Connection())
}

func InsertPasien(pasien model.Pasien) {
	t.ReadValue(pasien)
	sqlStr := "insert into " + t.TableName + " values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	tx, _ := pkg.Connection().Begin()
	stmt, _ := tx.Prepare(sqlStr)
	fmt.Println(t.Values)
	_, err := stmt.Exec(t.Values...)

	if err != nil {
		fmt.Println("miaww")
		panic(err)
	}
	tx.Commit()
}
