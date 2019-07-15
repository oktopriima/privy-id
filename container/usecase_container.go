/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:36 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package container

import (
	"github.com/oktopriima/privy-id/app/context/auth"
	"github.com/oktopriima/privy-id/app/context/product"
	"github.com/oktopriima/privy-id/app/context/role"
	"github.com/oktopriima/privy-id/app/context/roleuser"
	"github.com/oktopriima/privy-id/app/context/user"
	"go.uber.org/dig"
)

func BuildUsecaseProvider(container *dig.Container) *dig.Container {
	/** register your use case here with this format */
	var err error

	if err = container.Provide(auth.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(user.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(role.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(roleuser.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(product.NewUsecase); err != nil {
		panic(err)
	}
	return container
}
