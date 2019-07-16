/*
 * Copyright (c) 2019
 * Created at 7/16/19 6:25 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

type UpdateRequest interface {
	GetID() int64
	GetProductID() int64
	GetCategoryID() int64
}

type UpdateResponse interface {
}

func (uc *usecase) UpdateUsecase(request UpdateRequest) (UpdateResponse, error) {
	criteria := make(map[string]interface{})
	criteria["id"] = request.GetID()

	m, err := uc.productCategoryRepo.FindOneBy(criteria)
	if err != nil {
		return nil, err
	}

	m.CategoryID = request.GetCategoryID()
	m.ProductID = request.GetProductID()

	tx := uc.db.Begin()

	if err = uc.productCategoryRepo.Update(m, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return m, nil
}
