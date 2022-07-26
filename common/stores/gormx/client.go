package gormx

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Config struct {
		DSN string
	}

	DBManager struct {
		DB *gorm.DB
	}
)

func (m DBManager) GetDB(ctx context.Context) (*gorm.DB, error) {
	return m.DB.WithContext(ctx), nil
}

func MustBuildGormDB(conf Config) *DBManager {
	m, err := BuildDBManager(conf)
	if err != nil {
		logx.ErrorStack(err)
		panic(err)
	}

	return m
}

func BuildDBManager(conf Config) (*DBManager, error) { //, logger logger.Interface
	mysqlConfig := mysql.Config{
		DSN:                       conf.DSN, // DSN data source name
		DefaultStringSize:         191,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}
	client, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		//Logger:                                   logger, // 使用自定义 Logger
	})

	if err != nil {
		return nil, err
	}

	// TODO
	// db.SetMaxOpenConns(conf.MaxOpenConns) // 打开数据库连接的最大数量
	// db.SetMaxIdleConns(conf.MaxIdleConns) // 空闲连接池中连接的最大数量
	// db.SetConnMaxLifetime(conf.ConnMaxLifeTime)

	return &DBManager{
		DB: client,
	}, nil
}
