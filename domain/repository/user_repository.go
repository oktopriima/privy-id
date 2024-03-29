/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:27 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/domain/model"
)

type UserRepository interface {
	Create(users *model.Users, tx *gorm.DB) (*model.Users, error)
	Update(users *model.Users) (err error)
	Find(ID int64) (*model.Users, error)
	FindOneBy(map[string]interface{}) (*model.Users, error)
}
