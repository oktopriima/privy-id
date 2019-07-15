/*
 * Copyright (c) 2019
 * Created at 7/15/19 5:21 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

type updateRequest struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Enable      bool   `json:"enable" binding:"required"`
}

func (req updateRequest) GetID() int64 {
	return req.ID
}

func (req updateRequest) GetName() string {
	return req.Name
}

func (req updateRequest) GetDescription() string {
	return req.Description
}

func (req updateRequest) GetEnable() bool {
	return req.Enable
}
