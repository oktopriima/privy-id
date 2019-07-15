/*
 * Copyright (c) 2019
 * Created at 7/15/19 3:13 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

type createRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Enable      bool   `json:"enable" binding:"required"`
}

func (req createRequest) GetName() string {
	return req.Name
}

func (req createRequest) GetDescription() string {
	return req.Description
}

func (req createRequest) GetEnable() bool {
	return req.Enable
}
