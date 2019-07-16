/*
 * Copyright (c) 2019
 * Created at 7/16/19 7:45 PM
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

func (h *handler) DeleteHandler(ctx *gin.Context) {

	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := h.uc.DeleteUsecase(int64(ID))
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
