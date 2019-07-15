/*
 * Copyright (c) 2019
 * Created at 7/15/19 5:37 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package category

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/domain/repository"
)

type Usecase interface {
	CreateUsecase(request CreateRequest) (CreateResponse, error)
	UpdateUsecase(request UpdateRequest) (UpdateResponse, error)
	FindUsecase(ID int64) (FindResponse, error)
	FindAllUsecase() (FindAllResponse, error)
	DeleteUsecase(ID int64) (DeleteResponse, error)
}

type usecase struct {
	db           *gorm.DB
	categoryRepo repository.CategoryRepository
}

func NewUsecase(db *gorm.DB,
	categoryRepo repository.CategoryRepository) Usecase {
	return &usecase{db, categoryRepo}
}
