package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type config struct {
	appName       string
	appPort       int
	migrationPath string
	db            databaseConfig
	secretKey     string
	expiryTime    string
}

var appConfig config

func Load() {
	viper.SetDefault("APP_NAME", "database-connection")
	viper.SetDefault("APP_PORT", 8080)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	appConfig = config{
		appName:       readEnvString("APP_NAME"),
		appPort:       readEnvInt("APP_PORT"),
		migrationPath: readEnvString("MIGRATION_PATH"),
		db:            newDatabaseConfig(),
		secretKey:     readEnvString("SECRET_KEY"),
		expiryTime:    readEnvString("TOKEN_EXPIRATION_HOURS"),
	}
}

func AppName() string {
	return appConfig.appName
}

func AppPort() int {
	return appConfig.appPort
}

func MigrationPath() string {
	return appConfig.migrationPath
}

func readEnvInt(key string) int {
	checkIfSet(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid integer", key))
	}
	return v
}

func readEnvString(key string) string {
	checkIfSet(key)
	return viper.GetString(key)
}

func checkIfSet(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("Key %s is not set", key))
	}
}

func GetSecretKey() string {
	return appConfig.secretKey
}

func GetExpiryTime() string {
	return appConfig.expiryTime
}
