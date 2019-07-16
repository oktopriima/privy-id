/*
 * Copyright (c) 2019
 * Created at 7/15/19 6:47 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package model

type ProductCategory struct {
	ID         int64 `json:"id"`
	ProductID  int64 `json:"product_id"`
	CategoryID int64 `json:"category_id"`
}
