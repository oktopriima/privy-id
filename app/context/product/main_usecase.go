/*
 * Copyright (c) 2019
 * Created at 7/15/19 1:54 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/domain/repository"
)

type Usecase interface {
	CreateUsecase(request CreateRequest) (CreateResponse, error)
	UpdateUsecase(request UpdateRequest) (UpdateResponse, error)
	DeleteUsecase(request DeleteRequest) (DeleteResponse, error)
	FindAllUsecase() (FindAllResponse, error)
	FindUsecase(ID int64) (FindResponse, error)
}

type usecase struct {
	db          *gorm.DB
	productRepo repository.ProductRepository
}

func NewUsecase(db *gorm.DB,
	productRepo repository.ProductRepository) Usecase {
	return &usecase{db, productRepo}
}
