package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port                   string
	Host                   string
	Timeout                int
	DBDSN                  string
	DBTimeout              int
	JWTSecretKey           string
	AccessTokenExpiration  int
	RefreshTokenExpiration int
}

func NewConfig() *Config {
	port, err := getEnv("PORT")
	if err != nil {
		fmt.Println("Не удалось получить PORT из переменной окружения, используется порт по умолчанию")
	}
	host, err := getEnv("HOST")
	if err != nil {
		fmt.Println("Не удалось получить HOST из переменной окружения, используется хост по умолчанию")
	}
	timeout := 10
	if envValue, err := getEnv("SERVER_TIMEOUT"); err == nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			timeout = parsed
		}
	} else {
		fmt.Println("Не удалось получить SERVER_TIMEOUT из переменной окружения, используется 10 секунд")
	}
	dbTimeout := 5
	if envValue, err := getEnv("DB_TIMEOUT"); err == nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			dbTimeout = parsed
		}
	} else {
		fmt.Println("Не удалось получить DB_TIMEOUT из переменной окружения, используется 5 секунд")
	}

	dbDSN := "?"

	jwtSecretKey, err := getEnv("JWT_SECRET_KEY")
	if err != nil {
		fmt.Println("Не удалось получить JWT_SECRET_KEY из переменной окружения, используется значение по умолчанию")
	}
	accessTokenExpiration := 24
	if envValue, err := getEnv("JWT_ACCESS_TOKEN_EXPIRATION"); err == nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			accessTokenExpiration = parsed
		}
	}
	refreshTokenExpiration := 168
	if envValue, err := getEnv("JWT_REFRESH_TOKEN_EXPIRATION"); err == nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			refreshTokenExpiration = parsed
		}
	}

	return &Config{
		Port:                   port,
		Host:                   host,
		DBDSN:                  dbDSN,
		JWTSecretKey:           jwtSecretKey,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
		Timeout:                timeout,
		DBTimeout:              dbTimeout,
	}
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("s%: s%", "Не установлена переменная окружения", key)
	}
	return value, nil
}
