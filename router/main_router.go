/*
 * Copyright (c) 2019
 * Created at 5/20/19 11:02 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package router

import (
	"github.com/oktopriima/privy-id/app/config"
	"github.com/oktopriima/privy-id/httphandler/auth"
	"github.com/oktopriima/privy-id/httphandler/role"
	"github.com/oktopriima/privy-id/httphandler/roleuser"
	"github.com/oktopriima/privy-id/httphandler/user"
	"github.com/oktopriima/privy-id/libraries/middleware"
	"github.com/gin-gonic/gin"
)

func InvokeRoute(
	engine *gin.Engine,
	auth auth.Handler,
	user user.Handler,
	role role.Handler,
	roleuser roleuser.Handler,
) {

	engine.NoRoute()

	conf := config.NewConfig()

	markone := engine.Group("api/" + conf.GetString("app.version.tag") + conf.GetString("app.version.value"))
	markone.Use(gin.Logger())
	markone.Use(gin.Recovery())
	markone.Use(gin.ErrorLogger())

	/** auth route group */
	{
		authroute := markone.Group("auth")
		authroute.POST("/google", auth.GoogleLoginHandle)
		authroute.POST("/facebook", auth.FacebookLoginHandle)
		authroute.POST("/phone", auth.PhoneLoginHandler)
		authroute.POST("/email", auth.EmailLoginHandler)
	}

	/** profile route group */
	{
		me := markone.Group("me")
		me.Use(middleware.MyAuth())
		me.GET("", user.FindHandler)
	}

	/** role route group */
	{
		roleroute := markone.Group("roles")
		roleroute.GET("/generate", role.GenerateHandler)
	}

	/** roleuser route group */
	{
		roleuserroute := markone.Group("roleuser")
		roleuserroute.Use(middleware.MyAuth(config.ADMIN))
		roleuserroute.POST("", roleuser.CreateHandler)
		roleuserroute.DELETE(":id", roleuser.DeleteHandler)
	}

}
