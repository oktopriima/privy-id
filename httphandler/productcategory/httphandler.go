/*
 * Copyright (c) 2019
 * Created at 7/16/19 6:42 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/privy-id/app/context/productcategory"
)

type Handler interface {
	CreateHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
	FindHandler(ctx *gin.Context)
}

type handler struct {
	uc productcategory.Usecase
}

func NewHandler(uc productcategory.Usecase) Handler {
	return &handler{uc}
}
