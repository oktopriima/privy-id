/*
 * Copyright (c) 2019
 * Created at 7/15/19 5:31 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

type deleteRequest struct {
	ID int64 `json:"id"`
}

func (req deleteRequest) GetID() int64 {
	return req.ID
}
