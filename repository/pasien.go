package repository

import (
	"fmt"

	"../model"
	"../pkg"
)

func init() {
	fmt.Println("init table")
	var t pkg.Table
	t.ReadModel(model.Pasien{})
	t.Init(pkg.Connection())
	// pkg.InitTable(pkg.Connection(), pkg.ReadStructTags(model.Pasien{}))
}
