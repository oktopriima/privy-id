/*
 * Copyright (c) 2019
 * Created at 7/16/19 7:37 PM
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

func (h *handler) FindHandler(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := h.uc.FindUsecase(int64(productID))
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
