/*
 * Copyright (c) 2019
 * Created at 7/15/19 2:10 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

type UpdateRequest interface {
	GetID() int64
	GetName() string
	GetDescription() string
	GetEnable() bool
}

type UpdateResponse interface {
}

func (uc *usecase) UpdateUsecase(request UpdateRequest) (UpdateResponse, error) {
	m, err := uc.productRepo.Find(request.GetID())
	if err != nil {
		return nil, err
	}

	m.Enable = request.GetEnable()
	m.Description = request.GetDescription()
	m.Name = request.GetName()

	tx := uc.db.Begin()
	if err = uc.productRepo.Update(m, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return m, nil
}
