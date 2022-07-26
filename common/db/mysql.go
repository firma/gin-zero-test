package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"miya/common/config"
)

func GenerateMysqlGormDb(conf config.Database) (client *gorm.DB, err error) { //, logger logger.Interface
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s&parseTime=True&loc=Asia%%2FShanghai",
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.Database,
		conf.ConnTimeout,
		conf.ReadTimeout,
		conf.WriteTimeout)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	client, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		//Logger:                                   logger, // 使用自定义 Logger
	})

	if err != nil {
		return client, err
	}
	db, err := client.DB()
	if err != nil {
		return client, err
	}
	db.SetMaxOpenConns(conf.MaxOpenConns) // 打开数据库连接的最大数量
	db.SetMaxIdleConns(conf.MaxIdleConns) // 空闲连接池中连接的最大数量
	db.SetConnMaxLifetime(conf.ConnMaxLifeTime)

	return client, nil
}
