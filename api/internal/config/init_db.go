package config

import (
	"fmt"
	"gorm.io/gorm"
	"miya/common/db"
	"sync"
)

var (
	once   sync.Once
	Pools  map[string]*gorm.DB
	DbLink map[string]string
)

func GetDbInstance() {
	once.Do(func() {
		databases := GetAppConfig().Databases
		Pools = make(map[string]*gorm.DB)
		DbLink = make(map[string]string)
		for key, dbConfig := range databases {

			connection, err := db.GenerateMysqlGormDb(dbConfig)
			if err == nil {
				Pools[key] = connection
			}

		}
	})
}
func Connection(database ...interface{}) *gorm.DB {
	var databaseGroupName string
	if len(database) > 0 {
		databaseGroupName = database[0].(string)
	}
	fmt.Println(Pools)
	if _, isExists := Pools[databaseGroupName]; isExists == true {
		return Pools[databaseGroupName]
	}
	panic("database not found")
}
