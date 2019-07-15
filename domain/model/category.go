/*
 * Copyright (c) 2019
 * Created at 7/15/19 5:39 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package model

type Category struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
}
