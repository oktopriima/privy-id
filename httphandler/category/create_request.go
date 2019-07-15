/*
 * Copyright (c) 2019
 * Created at 7/15/19 6:14 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package category

type createRequest struct {
	Name   string `json:"name" binding:"required"`
	Enable bool   `json:"enable" binding:"required"`
}

func (req createRequest) GetName() string {
	return req.Name
}

func (req createRequest) GetEnable() bool {
	return req.Enable
}
