/*
 * Copyright (c) 2019
 * Created at 7/15/19 5:58 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package category

import "github.com/oktopriima/privy-id/domain/model"

type CreateRequest interface {
	GetName() string
	GetEnable() bool
}

type CreateResponse interface {
}

func (uc *usecase) CreateUsecase(request CreateRequest) (CreateResponse, error) {
	m := new(model.Category)
	m.Name = request.GetName()
	m.Enable = request.GetEnable()

	tx := uc.db.Begin()
	resp, err := uc.categoryRepo.Create(m, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return resp, nil
}
