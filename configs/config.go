package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost      string `mapstructure:"DB_host"`
	DBPort      string `mapstructure:"DB_port"`
	DBUser      string `mapstructure:"DB_user"`
	DBPassword  string `mapstructure:"DB_password"`
	DBName      string `mapstructure:"DB_name"`
	DBSslMode   string `mapstructure:"DB_sslmode"`
	DBSecretKey string `mapstructure:"SECRET_KEY"`
}

func LoadConfig(path string) (cfg *Config, err error) {

	cfg = new(Config)

	viper.SetConfigFile(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("Read config error %w", err)
		return
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		fmt.Errorf("Unmarshal config error %w", err)
		return
	}

	return cfg, nil
}
