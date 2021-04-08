package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHttpPort               = "8000"
	defaultHttpRWTimeout          = 10 * time.Second
	defaultHttpMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
	defaultLimiterRPS             = 10
	defaultLimiterBurst           = 2
	defaultLimiterTTL             = 10 * time.Minute
	defaultVerificationCodeLength = 8
)

type (
	Config struct {
		Postgres PostgresConfig
		HTTP     HTTPConfig
		Auth     AuthConfig
		CacheTTL time.Duration `mapstructure:"ttl"`
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
	PostgresConfig struct {
		DBname   string
		User     string
		Password string
		Sslmode  string
	}
	AuthConfig struct {
		JWT                    JWTConfig
		PasswordSalt           string
		VerificationCodeLength int `mapstructure:"verificationCodeLength"`
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
	}
)

func Init(path string) (*Config, error) {
	populateDefaults()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	if err := parseEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)
	fmt.Println(cfg.Postgres)
	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("cache.ttl", &cfg.CacheTTL); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("auth", &cfg.Auth.JWT); err != nil {
		return err
	}
	return nil
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // config file name

	return viper.ReadInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultHttpMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)
	viper.SetDefault("auth.accessTokenTTL", defaultAccessTokenTTL)
	viper.SetDefault("auth.refreshTokenTTL", defaultRefreshTokenTTL)
	viper.SetDefault("auth.verificationCodeLength", defaultVerificationCodeLength)
	viper.SetDefault("limiter.rps", defaultLimiterRPS)
	viper.SetDefault("limiter.burst", defaultLimiterBurst)
	viper.SetDefault("limiter.ttl", defaultLimiterTTL)
}

func parseEnv() error {
	if err := parsePostgresEnvVariables(); err != nil {
		return err
	}
	if err := parseJWTFromEnv(); err != nil {
		return err
	}
	return parsePasswordFromEnv()
}
func setFromEnv(cfg *Config) {
	cfg.Postgres.DBname = viper.GetString("dbname")
	cfg.Postgres.User = viper.GetString("user")
	cfg.Postgres.Password = viper.GetString("password")
	cfg.Postgres.Sslmode = viper.GetString("sslmode")
	cfg.Auth.PasswordSalt = viper.GetString("salt")
	cfg.Auth.JWT.SigningKey = viper.GetString("signingkey")

	fmt.Println(cfg.Postgres)

}
func parsePostgresEnvVariables() error {

	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_DBNAME", "hub")

	os.Setenv("POSTGRES_PASSWORD", "123")

	os.Setenv("POSTGRES_SSLMODE", "disabled")

	viper.SetEnvPrefix("postgres")
	if err := viper.BindEnv("user"); err != nil {
		return err
	}

	if err := viper.BindEnv("dbname"); err != nil {
		return err
	}
	if err := viper.BindEnv("password"); err != nil {
		return err
	}
	fmt.Println(os.Getenv("POSTGRES_PASSWORD"), "111111111111")
	return viper.BindEnv("sslmode")

}
func parsePasswordFromEnv() error {
	os.Setenv("PASSWORD_SALT", "1234")

	viper.SetEnvPrefix("password")
	return viper.BindEnv("salt")
}
func parseJWTFromEnv() error {
	os.Setenv("JWT_SIGNINGKEY", "signing_key")

	viper.SetEnvPrefix("jwt")
	return viper.BindEnv("signingkey")
}
