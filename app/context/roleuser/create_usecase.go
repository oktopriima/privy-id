/*
 * Copyright (c) 2019
 * Created at 5/29/19 10:33 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package roleuser

import "github.com/oktopriima/privy-id/domain/model"

type CreateRequest interface {
	GetUserID() int64
	GetRoleID() int64
}

type CreateResponse interface {
	GetData() *model.UserRole
}

func (uc *usecase) Create(request CreateRequest) (CreateResponse, error) {
	tx := uc.db.Begin()
	m := new(model.UserRole)
	m.UsersID = request.GetUserID()
	m.RolesID = request.GetRoleID()

	resp, err := uc.roleUserRepo.Create(m, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return createResponse{
		Data: resp,
	}, nil
}
