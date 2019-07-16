/*
 * Copyright (c) 2019
 * Created at 7/15/19 7:04 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/domain/repository"
)

type Usecase interface {
	CreateUsecase(request CreateRequest) (CreateResponse, error)
	FindUsecase(productID int64) (FindResponse, error)
	UpdateUsecase(request UpdateRequest) (UpdateResponse, error)
	DeleteUsecase(ID int64) (DeleteResponse, error)
}

type usecase struct {
	db                  *gorm.DB
	productCategoryRepo repository.ProductCategoryRepository
}

func NewUsecase(db *gorm.DB,
	productCategoryRepo repository.ProductCategoryRepository) Usecase {
	return &usecase{db, productCategoryRepo}
}
