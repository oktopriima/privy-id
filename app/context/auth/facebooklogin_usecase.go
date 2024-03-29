/*
 * Copyright (c) 2019
 * Created at 5/28/19 10:27 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/app/config"
	"github.com/oktopriima/privy-id/domain/model"
	"github.com/oktopriima/privy-id/libraries/helper"
	"github.com/oktopriima/privy-id/libraries/middleware"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) FacebookLogin(request AuthRequest) (AuthResponse, error) {
	fResp, err := helper.GetFacebookData(request.GetSocialToken())
	if err != nil {
		return nil, err
	}

	/** get user */
	user, err := uc.getUserByEmail(fResp.Email)
	if err != nil || user.ID == 0 {
		tx := uc.db.Begin()
		user, err := uc.registerFromFacebook(fResp, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if _, err = uc.createSocialMedia(user, fResp.ID, request, tx); err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
	}

	social, err := uc.getUserSocial(fResp.ID)
	if err != nil {
		return nil, err
	}

	if social.UsersID != user.ID {
		return nil, errors.New("user didn't match")
	}

	var role string
	roles, err := uc.findRoleUser(user)
	if err != nil {
		role = config.CONSUMER
	} else {
		role = roles.Value
	}

	/** generate parameter for Custom JWT */
	param := middleware.TokenStructure{}
	param.UserID = user.ID
	param.Email = user.Email
	param.Role = role

	auth := middleware.NewCustomAuth([]byte(config.SIGNATURE))
	token, err := auth.GenerateToken(param)
	if err != nil {
		return nil, err
	}

	if err = uc.updateLastLogin(user); err != nil {
		return nil, err
	}

	return token, nil
}

func (uc *usecase) registerFromFacebook(response *helper.FacebookResponse, tx *gorm.DB) (*model.Users, error) {
	pass, _ := bcrypt.GenerateFromPassword([]byte(helper.RandString(10)), bcrypt.DefaultCost)

	m := new(model.Users)

	m.Name = response.Name
	m.Username = response.Email
	m.Email = response.Email
	m.Phone = ""
	m.Password = string(pass)
	m.LastLogin = time.Now()
	m.SecondaryEmail = response.Email
	m.IsVerified = false
	m.Avatar = response.Avatar

	if m, err := uc.userRepo.Create(m, tx); err != nil {
		return nil, err
	} else {
		return m, nil
	}
}
