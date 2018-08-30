package sconfig

import "github.com/spf13/viper"

func JuheKey() string {
	return viper.GetString("juheKey")
}

func HttpPort() string {
	return viper.GetString("listen")
}

func WechatCorID() string {
	return viper.GetString("wechatCorID")
}

func WechatAgentID() int64 {
	return viper.GetInt64("wechatAgentID")
}

func WechatSecret() string {
	return viper.GetString("wechatSecret")
}
