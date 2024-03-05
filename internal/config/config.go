package config

import "github.com/pelletier/go-toml"

var Bot struct {
	Debug    bool   `toml:"debug"`
	Timeout  int    `toml:"timeout"`
	BotToken string `toml:"bot_token"`
	GroupId  int64  `toml:"group_id"`
}

var Juejin struct {
	Uuid      string `toml:"uuid"`
	SessionId string `toml:"session_id"`
}

var configFile = "config.toml"

func init() {
	conf, err := toml.LoadFile(configFile)
	if err != nil {
		panic("Failed to open config file")
	}
	if err := conf.Get("bot").(*toml.Tree).Unmarshal(&Bot); err != nil {
		panic("mapping [bot] section")
	}

	if err = conf.Get("juejin").(*toml.Tree).Unmarshal(&Juejin); err != nil {
		panic("mapping [juejin] section")
	}
}
