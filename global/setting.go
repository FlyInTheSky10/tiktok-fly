package global

import (
	"github.com/FlyInThesky10/TikTok-Fly/pkg/logger"
	"github.com/FlyInThesky10/TikTok-Fly/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
