package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DSN    string        `mapstructure:"dsn"`
	Server *ConfigServer `mapstructure:"server"`
	SMTP   *ConfigSMTP   `mapstructure:"smtp"`
}

type ConfigServer struct {
	Addr string `mapstructure:"addr"`
}

type ConfigSMTP struct {
	Hostname string `mapstructure:"hostname"`
	Nicename string `mapstructure:"nicename"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func New() *Config {
	return &Config{
		DSN: "root:password@tcp(localhost:3306)/test?charset=utf8mb4&collation=utf8mb4_unicode_ci&&parseTime=True&loc=Local",
		Server: &ConfigServer{
			Addr: "localhost:8080",
		},
		SMTP: &ConfigSMTP{
			Hostname: "",
			Nicename: "",
			Username: "",
			Password: "",
		},
	}
}

func (cfg *Config) Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.xjtuana-api/")
	viper.AddConfigPath("$HOME/.config/xjtuana-api/")
	viper.AddConfigPath("/etc/xjtuana-api/")
	if err := viper.ReadInConfig(); err != nil {
		panic("failed to read config file, error: " + err.Error())
	}
	tmp := &Config{}
	if err := viper.UnmarshalKey("api.micro.mail", tmp); err != nil {
		panic("failed to init api config, error: " + err.Error())
	}
	if tmp.DSN != "" {
		cfg.DSN = tmp.DSN
	}
	if tmp.Server != nil {
		if tmp.Server.Addr != "" {
			cfg.Server.Addr = tmp.Server.Addr
		}
	}
	if tmp.SMTP != nil {
		if tmp.SMTP.Hostname != "" {
			cfg.SMTP.Hostname = tmp.SMTP.Hostname
		}
		if tmp.SMTP.Nicename != "" {
			cfg.SMTP.Nicename = tmp.SMTP.Nicename
		}
		if tmp.SMTP.Username != "" {
			cfg.SMTP.Username = tmp.SMTP.Username
		}
		if tmp.SMTP.Password != "" {
			cfg.SMTP.Password = tmp.SMTP.Password
		}
	}
}
