/*
 * Copyright (c) 2019
 * Created at 7/15/19 6:06 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package category

type UpdateRequest interface {
	GetID() int64
	GetName() string
	GetEnable() bool
}

type UpdateResponse interface {
}

func (uc *usecase) UpdateUsecase(request UpdateRequest) (UpdateResponse, error) {
	m, err := uc.categoryRepo.Find(request.GetID())
	if err != nil {
		return nil, err
	}

	m.Name = request.GetName()
	m.Enable = request.GetEnable()

	tx := uc.db.Begin()
	if err = uc.categoryRepo.Update(m, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return m, nil
}
