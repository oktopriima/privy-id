/*
 * Copyright (c) 2019
 * Created at 7/15/19 7:06 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

import "github.com/oktopriima/privy-id/domain/model"

type CreateRequest interface {
	GetProductID() int64
	GetCategoryID() int64
}

type CreateResponse interface {
}

func (uc *usecase) CreateUsecase(request CreateRequest) (CreateResponse, error) {
	m := new(model.ProductCategory)
	m.ProductID = request.GetProductID()
	m.CategoryID = request.GetCategoryID()

	tx := uc.db.Begin()
	resp, err := uc.productCategoryRepo.Create(m, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return resp, nil
}
