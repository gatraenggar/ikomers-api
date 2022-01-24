package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DbTestHost string
	DbTestUser string
	DbTestPass string
	DbTestPort string
	DbTestName string

	AccessTokenKey  string
	RefreshTokenKey string
}

func GetConfig() *Config {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./../..")
	viper.AddConfigPath("./../../..")
	viper.AddConfigPath("./../../../..")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	dbTestHost, ok := viper.Get("DBTEST_HOST").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	dbTestPort, ok := viper.Get("DBTEST_PORT").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	dbTestUser, ok := viper.Get("DBTEST_USER").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	dbTestPass, ok := viper.Get("DBTEST_PASS").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	dbTestName, ok := viper.Get("DBTEST_NAME").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	accessTokenKey, ok := viper.Get("ACCESS_TOKEN_KEY").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	refreshTokenKey, ok := viper.Get("REFRESH_TOKEN_KEY").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return &Config{
		DbTestHost:      dbTestHost,
		DbTestPort:      dbTestPort,
		DbTestUser:      dbTestUser,
		DbTestPass:      dbTestPass,
		DbTestName:      dbTestName,
		AccessTokenKey:  accessTokenKey,
		RefreshTokenKey: refreshTokenKey,
	}
}
