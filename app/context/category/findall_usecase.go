/*
 * Copyright (c) 2019
 * Created at 7/15/19 6:05 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package category

type FindAllResponse interface {
}

func (uc *usecase) FindAllUsecase() (FindAllResponse, error) {
	resp, err := uc.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
