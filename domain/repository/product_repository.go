package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/domain/model"
)

type ProductRepository interface {
	Create(product *model.Product, tx *gorm.DB) (*model.Product, error)
	Find(ID int64) (*model.Product, error)
	FindAll() ([]*model.Product, error)
	Update(product *model.Product, tx *gorm.DB) (err error)
	Delete(user *model.Product, tx *gorm.DB) (err error)
}
