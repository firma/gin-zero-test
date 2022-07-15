package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/gen"
)

type EmptyModel struct {
	ID int32
}

// IsEmpty 判断空
func (m *EmptyModel) IsEmpty() bool {
	if m == nil {
		return true
	}

	return m.ID == 0
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./api/internal/query",
	})

	db, err := gorm.Open(mysql.Open("root:root@(127.0.0.1:13306)/dao_gen_db?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Panicf("connect to db failed")
	}
	g.UseDB(db)

	isEmptyMethod := (&EmptyModel{}).IsEmpty

	var (
		createTimeOpt = gen.FieldGORMTag("create_time", "column:create_time;autoCreateTime")
		updateTimeOpt = gen.FieldGORMTag("update_time", "column:update_time;autoUpdateTime")
		modifyTimeOpt = gen.FieldGORMTag("modify_time", "column:modify_time;autoUpdateTime")
		deleteTimeOpt = gen.FieldType("delete_time", "soft_delete.DeletedAt")
	)
	opts := []gen.FieldOpt{createTimeOpt, updateTimeOpt, modifyTimeOpt, deleteTimeOpt}

	user := g.GenerateModelAs("user", "User", opts...).AddMethod(isEmptyMethod)

	g.ApplyBasic(
		user,
	)

	g.Execute()
}
