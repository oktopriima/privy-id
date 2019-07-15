/*
 * Copyright (c) 2019
 * Created at 7/15/19 2:19 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

type DeleteRequest interface {
	GetID() int64
}

type DeleteResponse interface {
}

func (uc *usecase) DeleteUsecase(request DeleteRequest) (DeleteResponse, error) {
	m, err := uc.productRepo.Find(request.GetID())
	if err != nil {
		return nil, err
	}

	tx := uc.db.Begin()
	if err = uc.productRepo.Delete(m, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return m, nil
}
