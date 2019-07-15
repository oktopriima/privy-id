/*
 * Copyright (c) 2019
 * Created at 7/15/19 1:54 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

import "github.com/oktopriima/privy-id/domain/model"

type CreateRequest interface {
	GetName() string
	GetDescription() string
	GetEnable() bool
}

type CreateResponse interface {
}

func (uc *usecase) CreateUsecase(request CreateRequest) (CreateResponse, error) {
	m := new(model.Product)
	m.Name = request.GetName()
	m.Description = request.GetDescription()
	m.Enable = request.GetEnable()

	tx := uc.db.Begin()

	resp, err := uc.productRepo.Create(m, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return resp, nil
}
