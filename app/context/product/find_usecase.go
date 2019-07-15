/*
 * Copyright (c) 2019
 * Created at 7/15/19 2:08 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

type FindResponse interface {
}

func (uc *usecase) FindUsecase(ID int64) (FindResponse, error) {

	resp, err := uc.productRepo.Find(ID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
