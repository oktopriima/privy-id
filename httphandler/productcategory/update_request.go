/*
 * Copyright (c) 2019
 * Created at 7/16/19 7:40 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

type updateRequest struct {
	ID         int64 `json:"id"`
	ProductID  int64 `json:"product_id"`
	CategoryID int64 `json:"category_id"`
}

func (req updateRequest) GetID() int64 {
	return req.ID
}

func (req updateRequest) GetProductID() int64 {
	return req.ProductID
}

func (req updateRequest) GetCategoryID() int64 {
	return req.CategoryID
}
