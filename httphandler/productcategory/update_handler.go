/*
 * Copyright (c) 2019
 * Created at 7/16/19 7:41 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package productcategory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oktopriima/privy-id/libraries/httpresponse"
)

func (h *handler) UpdateHandler(ctx *gin.Context) {
	var request updateRequest

	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	request.ID = int64(ID)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := h.uc.UpdateUsecase(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
