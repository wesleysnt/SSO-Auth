package repositories

import (
	"sso-auth/app/facades"

	"gorm.io/gorm"
)

type TxRepository struct {
	ormTx *gorm.DB
}

func NewTxRepository() *TxRepository {
	return &TxRepository{ormTx: facades.Orm()}
}

func (r *TxRepository) Begin() *gorm.DB {
	return r.ormTx.Begin()
}

func (r *TxRepository) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (r *TxRepository) Rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}
