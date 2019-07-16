/*
 * Copyright (c) 2019
 * Created at 7/16/19 6:45 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

type createRequest struct {
	ProductID  int64 `json:"product_id"`
	CategoryID int64 `json:"category_id"`
}

func (req createRequest) GetProductID() int64 {
	return req.ProductID
}

func (req createRequest) GetCategoryID() int64 {
	return req.CategoryID
}
