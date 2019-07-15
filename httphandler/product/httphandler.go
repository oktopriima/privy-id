/*
 * Copyright (c) 2019
 * Created at 7/15/19 3:11 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/privy-id/app/context/product"
)

type Handler interface {
	CreateHandler(ctx *gin.Context)
	FindHandler(ctx *gin.Context)
	FindAllHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
}

type handler struct {
	uc product.Usecase
}

func NewHandler(uc product.Usecase) Handler {
	return &handler{uc}
}
