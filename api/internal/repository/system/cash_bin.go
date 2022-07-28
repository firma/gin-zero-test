package system

import (
	"context"
	"gorm.io/gorm"
	"miya/api/internal/query"
)

type (
	ICasBinRepo interface {
		GetDb(ctx context.Context) *gorm.DB
	}

	CasBinRepo struct {
	}
)

func (c CasBinRepo) GetDb(ctx context.Context) *gorm.DB {
	//TODO implement me
	return query.CasBin(ctx)
}

func NewCasBinRepo() ICasBinRepo {
	return CasBinRepo{}
}
