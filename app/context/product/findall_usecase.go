/*
 * Copyright (c) 2019
 * Created at 7/15/19 3:06 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

type FindAllResponse interface {
}

func (uc *usecase) FindAllUsecase() (FindAllResponse, error) {
	data, err := uc.productRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}
