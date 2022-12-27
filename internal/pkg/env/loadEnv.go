package env

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost            string `mapstructure:"POSTGRES_HOST"`
	DBUserName        string `mapstructure:"POSTGRES_USER"`
	DBUserPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	DBName            string `mapstructure:"POSTGRES_DB"`
	DBPort            string `mapstructure:"POSTGRES_PORT"`
	ServerPort        string `mapstructure:"PORT"`
	ClientOrigin      string `mapstructure:"CLIENT_ORIGIN"`
	ETHERSCAN_API_KEY string `mapstructure:"ETHERSCAN_API_KEY"`
	CONTRACT_ADDRESS  string `mapstructure:"CONTRACT_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Infof("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&config)
	return config, err
}
