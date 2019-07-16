/*
 * Copyright (c) 2019
 * Created at 7/15/19 7:09 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

type FindResponse interface {
}

func (uc *usecase) FindUsecase(productID int64) (FindResponse, error) {
	criteria := make(map[string]interface{})
	criteria["product_id"] = productID

	resp, err := uc.productCategoryRepo.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
