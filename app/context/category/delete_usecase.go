/*
 * Copyright (c) 2019
 * Created at 7/15/19 6:09 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package category

type DeleteResponse interface {
}

func (uc *usecase) DeleteUsecase(ID int64) (DeleteResponse, error) {
	m, err := uc.categoryRepo.Find(ID)
	if err != nil {
		return nil, err
	}

	tx := uc.db.Begin()
	if err = uc.categoryRepo.Delete(m, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return m, nil
}
