/*
 * Copyright (c) 2019
 * Created at 7/15/19 6:12 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package category

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/privy-id/app/context/category"
)

type Handler interface {
	CreateHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	FindHandler(ctx *gin.Context)
	FindAllHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
}

type handler struct {
	uc category.Usecase
}

func NewHandler(uc category.Usecase) Handler {
	return &handler{uc}
}
