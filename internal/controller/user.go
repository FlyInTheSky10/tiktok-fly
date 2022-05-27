package controller

import (
	"github.com/FlyInThesky10/TikTok-Fly/global"
	"github.com/FlyInThesky10/TikTok-Fly/internal/service"
	"github.com/FlyInThesky10/TikTok-Fly/internal/service/param"
	"github.com/FlyInThesky10/TikTok-Fly/pkg/app"
	"github.com/FlyInThesky10/TikTok-Fly/pkg/errcode"
	"github.com/FlyInThesky10/TikTok-Fly/pkg/util"
	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	req := param.ReqUserRegister{}
	if valid, errs := app.BindAndValid(ctx, &req); !valid {
		global.Logger.Errorf(ctx, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(ctx)
	userId, err := svc.UserRegister(req.Username, req.Password)
	if err != nil {
		global.Logger.Errorf(ctx, "svc.UserRegister err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterError)
		return
	}

	token, _, err := app.GenerateJWTToken(req.Username, util.EncodeMD5(req.Password), userId)
	if err != nil {
		global.Logger.Errorf(ctx, "app.GenerateJWTToken err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterError)
		return
	}

	response.ToResponse(param.ResUserRegister{
		Response: param.Response{StatusCode: 0, StatusMsg: "success"},
		Token:    token,
		UserId:   userId,
	})
}
func UserLogin(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	req := param.ReqUserLogin{}
	if valid, errs := app.BindAndValid(ctx, &req); !valid {
		global.Logger.Errorf(ctx, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(ctx)
	token, userId, err := svc.UserLogin(req.Username, req.Password)
	if err != nil {
		global.Logger.Errorf(ctx, "svc.UserLogin err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterError)
		return
	}

	response.ToResponse(param.ResUserLogin{
		Response: param.Response{StatusCode: 0, StatusMsg: "success"},
		Token:    token,
		UserId:   userId,
	})
}
func UserInfo(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	req := param.ReqUserInfo{}
	if valid, errs := app.BindAndValid(ctx, &req); !valid {
		global.Logger.Errorf(ctx, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(ctx)
	claims, err := app.ParseJWTToken(req.Token)
	if err != nil {
		global.Logger.Errorf(ctx, "app.ParseJWTToken errs: %v", err)
		response.ToErrorResponse(err)
		return
	}

	userInfo, err2 := svc.UserInfo(req.UserId, claims.UserId)
	if err2 != nil {
		global.Logger.Errorf(ctx, "svc.UserInfo errs: %v", err2)
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	response.ToResponse(userInfo)
}
