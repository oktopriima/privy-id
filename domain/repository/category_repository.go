/*
 * Copyright (c) 2019
 * Created at 7/15/19 5:39 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/domain/model"
)

type CategoryRepository interface {
	Create(category *model.Category, tx *gorm.DB) (*model.Category, error)
	Find(ID int64) (*model.Category, error)
	FindAll() ([]*model.Category, error)
	Update(category *model.Category, tx *gorm.DB) (err error)
	Delete(category *model.Category, tx *gorm.DB) (err error)
}
