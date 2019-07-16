/*
 * Copyright (c) 2019
 * Created at 7/16/19 6:33 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

type DeleteResponse interface {
}

func (uc *usecase) DeleteUsecase(ID int64) (DeleteResponse, error) {
	criteria := make(map[string]interface{})
	criteria["id"] = ID

	m, err := uc.productCategoryRepo.FindOneBy(criteria)
	if err != nil {
		return nil, err
	}

	tx := uc.db.Begin()

	if err = uc.productCategoryRepo.Delete(m, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return m, nil
}
