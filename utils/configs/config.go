package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            int `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WEBServerHost     string `mapstructure:"WEB_SERVER_HOST"`
	WEBServerPort     int `mapstructure:"WEB_SERVER_PORT"`
	WEBServerTimeOut  int64  `mapstructure:"GRPC_SERVER_TIMEOUT"`
	GRPCServerHost    string `mapstructure:"GRPC_SERVER_HOST"`
	GRPCServerPort    int `mapstructure:"GRPC_SERVER_PORT"`
	GRPCServerTimeOut int64  `mapstructure:"GRPC_SERVER_TIMEOUT"`
	// JWTSecret      string `mapstructure:"JWT_SECRET"`
	// JWTExpiresIn   int    `mapstructure:"JWT_EXPIRES_IN"`
	// TokenAuth      *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("main_app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	// log.Println("JWTSecret", cfg.JWTSecret)
	// cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
