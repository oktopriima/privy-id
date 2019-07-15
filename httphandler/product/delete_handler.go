/*
 * Copyright (c) 2019
 * Created at 7/15/19 5:30 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oktopriima/privy-id/libraries/httpresponse"
)

func (h *handler) DeleteHandler(ctx *gin.Context) {
	var request deleteRequest

	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	request.ID = int64(ID)

	resp, err := h.uc.DeleteUsecase(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
