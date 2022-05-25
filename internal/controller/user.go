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
	err := svc.UserRegister(req.Username, req.Password)
	if err != nil {
		global.Logger.Errorf(ctx, "svc.UserRegister err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterError)
		return
	}

	token, _, err := app.GenerateJWTToken(req.Username, util.EncodeMD5(req.Password))
	if err != nil {
		global.Logger.Errorf(ctx, "app.GenerateJWTToken err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterError)
		return
	}

	response.ToResponse(param.ResUserRegister{Token: token})
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
	token, err := svc.UserLogin(req.Username, req.Password)
	if err != nil {
		global.Logger.Errorf(ctx, "svc.UserLogin err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterError)
		return
	}

	response.ToResponse(param.ResUserLogin{Token: token})
}
