package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Database *ConfigDatabase         `yaml:"database"`
	Imapc    *ConfigImapConcentrator `yaml:"imapc"`
	Poller   *ConfigPoller           `yaml:"poller"`
}

type ConfigDatabase struct {
	Host     string `yaml:"host" env:"DBHOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DBPORT" env-default:"5432"`
	Name     string `yaml:"name" env:"DBNAME" env-default:"postgres"`
	User     string `yaml:"user" env:"DBUSER" env-default:"root"`
	Password string `yaml:"password" env:"DBPASSWORD"`
}

type ConfigImapConcentrator struct {
	Addr      string    `yaml:"addr" env:"ADDR" env-default:"localhost:7070"`
	Whitelist *[]string `yaml:"whitelist" env:"WHITELIST"`
}

type ConfigPoller struct {
	Cron string `yaml:"cron" env:"CRON" env-default:"*/10 * * * *"`
}

func ReadConfig(filename string) (*Config, error) {
	cfg := Config{}
	if err := cleanenv.ReadConfig(filename, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
