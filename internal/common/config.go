package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config struct to hold configuration values
type Config struct {
	PostgresHost string
	PostgresPort int
	RedisHost    string
	RedisPort    int
}

// LoadConfig reads configuration from environment variables
func LoadConfig() (config Config) {
	viper.SetConfigFile(".env")

	er := viper.ReadInConfig()

	if er != nil {
		log.Fatalln(er)
	}
	viper.AutomaticEnv()

	// Read environment variables
	viper.BindEnv("PostgresHost")
	viper.BindEnv("PostgresPort")
	viper.BindEnv("RedisHost")
	viper.BindEnv("RedisPort")

	// Populate config struct
	config.PostgresHost = viper.GetString("PostgresHost")
	config.PostgresPort = viper.GetInt("PostgresPort")
	config.RedisHost = viper.GetString("RedisHost")
	config.RedisPort = viper.GetInt("RedisPort")

	return
}
