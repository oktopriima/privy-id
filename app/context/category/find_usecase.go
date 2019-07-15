/*
 * Copyright (c) 2019
 * Created at 7/15/19 6:03 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package category

type FindResponse interface {
}

func (uc *usecase) FindUsecase(ID int64) (FindResponse, error) {
	resp, err := uc.categoryRepo.Find(ID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
