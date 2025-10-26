package config

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
