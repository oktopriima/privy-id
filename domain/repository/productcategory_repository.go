/*
 * Copyright (c) 2019
 * Created at 7/15/19 6:49 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/domain/model"
)

type ProductCategoryRepository interface {
	Create(category *model.ProductCategory, tx *gorm.DB) (*model.ProductCategory, error)
	FindBy(criteria map[string]interface{}) ([]*model.ProductCategory, error)
	Update(category *model.ProductCategory, tx *gorm.DB) (err error)
	Delete(category *model.ProductCategory, tx *gorm.DB) (err error)
	FindOneBy(criteria map[string]interface{}) (*model.ProductCategory, error)
}
