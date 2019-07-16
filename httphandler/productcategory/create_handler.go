/*
 * Copyright (c) 2019
 * Created at 7/16/19 6:46 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktopriima/privy-id/libraries/httpresponse"
)

func (h *handler) CreateHandler(ctx *gin.Context) {
	var request createRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := h.uc.CreateUsecase(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
