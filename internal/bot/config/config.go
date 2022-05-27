package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Imapc ConfigImapConcentrator `yaml:"imapc"`
	TgBot TelegramBot            `yaml:"telegram"`
}

type ConfigImapConcentrator struct {
	Addr string `yaml:"addr" env:"ICADDR" env-default:"127.0.0.1:7070"`
	Key  string `yaml:"key" env:"ICKEY"`
}

type TelegramBot struct {
	Token string `yaml:"token" env:"TGTOKEN"`
}

func ReadConfig(filename string) (*Config, error) {
	cfg := Config{}
	if err := cleanenv.ReadConfig(filename, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
