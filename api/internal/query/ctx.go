package query

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"gorm.io/gorm"
)

type (
	ctxKey struct{}

	DBManager interface {
		GetDB(ctx context.Context) (*gorm.DB, error)
	}
)

var (
	dbManager DBManager
)

// SetUpDBManager 需初始化DB manager
func SetUpDBManager(manager DBManager) {
	dbManager = manager
}

func Ctx(ctx context.Context) *Query {
	if iface := ctx.Value(ctxKey{}); iface != nil {
		q, ok := iface.(*Query)
		if !ok {
			log.Panicf("unexpect context value type: %s", reflect.TypeOf(q))
		}

		return q
	}

	db, err := dbManager.GetDB(ctx)
	if err != nil {
		log.Panicf("获取DB client err: %s", err)
	}

	return Use(db)
}
func CasBin(ctx context.Context) *gorm.DB {
	db, err := dbManager.GetDB(ctx)
	if err != nil {
		log.Panicf("获取DB client err: %s", err)
	}

	return db
}

func Transaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	if dbManager == nil {
		return fmt.Errorf("db manager is nil")
	}

	db, err := dbManager.GetDB(ctx)
	if err != nil {
		return err
	}

	return Use(db).Transaction(func(db *Query) error {
		txCtx := ctxWithTransaction(ctx, db)

		return fn(txCtx)
	})
}

func ctxWithTransaction(ctx context.Context, q *Query) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, ctxKey{}, q)
}
