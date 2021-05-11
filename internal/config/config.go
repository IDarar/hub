package config

import (
	"strings"
	"time"

	"github.com/IDarar/hub/pkg/logger"
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
		Redis    RedisConfig
		HTTP     HTTPConfig
		Auth     AuthConfig
		CacheTTL time.Duration `mapstructure:"ttl"`
		TestPostgresConfig
		GRPC GRPCConfig
	}
	GRPCConfig struct {
		Port             string `mapstructure:"port"`
		ServerCertFile   string `mapstructure:"servercertfile"`
		ServerKeyFile    string `mapstructure:"serverkeyfile"`
		ClientCACertFile string `mapstructure:"clientcacertfile"`
		ClientKeyFile    string `mapstructure:"clientkeyfile"`
		ClientCertFile   string `mapstructure:"clientcertfile"`
	}
	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
	PostgresConfig struct {
		Host     string `mapstructure:"POSTGRES_HOST"`
		Port     string `mapstructure:"POSTGRES_PORT"`
		DBname   string `mapstructure:"POSTGRES_DBNAME"`
		User     string `mapstructure:"POSTGRES_USER"`
		Password string `mapstructure:"POSTGRES_PASSWORD"`
	}
	RedisConfig struct {
		Addr     string `mapstructure:"REDIS_ADDR"`
		Password string `mapstructure:"REDIS_PASSWORD"`
		DB       int    `mapstructure:"REDIS_DB"`
	}
	TestPostgresConfig struct {
		URL string
	}
	AuthConfig struct {
		JWT                    JWTConfig
		PasswordSalt           string `mapstructure:"PASSWORD_SALT"`
		VerificationCodeLength int    `mapstructure:"verificationCodeLength"`
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string        `mapstructure:"JWT_SIGNINGKEY"`
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
	logger.Info("Postgres ", cfg.Postgres)
	logger.Info("Redis ", cfg.Redis)
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
	if err := viper.UnmarshalKey("grpc", &cfg.GRPC); err != nil {
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
	if err := parseRedisEnvVariables(); err != nil {
		return err
	}
	if err := parseJWTFromEnv(); err != nil {
		return err
	}
	return parsePasswordFromEnv()
}

//TODO
//Are set in docker compose with .env
func setFromEnv(cfg *Config) {
	cfg.Postgres.Host = viper.GetString("host")
	cfg.Postgres.Port = viper.GetString("port")
	cfg.Postgres.DBname = viper.GetString("dbname")
	cfg.Postgres.User = viper.GetString("user")
	cfg.Postgres.Password = viper.GetString("password")
	cfg.Redis.DB = viper.GetInt("db")
	cfg.Redis.Addr = viper.GetString("addr")
	cfg.Redis.Password = viper.GetString("redispassword")

	cfg.TestPostgresConfig.URL = viper.GetString("url")

	cfg.Auth.PasswordSalt = viper.GetString("salt")
	cfg.Auth.JWT.SigningKey = viper.GetString("signingkey")

}
func parsePostgresEnvVariables() error {

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
	if err := viper.BindEnv("port"); err != nil {
		return err
	}
	if err := viper.BindEnv("host"); err != nil {
		return err
	}

	viper.SetEnvPrefix("database")
	if err := viper.BindEnv("url"); err != nil {
		return err
	}
	return nil

}
func parseRedisEnvVariables() error {

	viper.SetEnvPrefix("redis")
	if err := viper.BindEnv("addr"); err != nil {
		return err
	}

	if err := viper.BindEnv("redispassword"); err != nil {
		return err
	}
	if err := viper.BindEnv("db"); err != nil {
		return err
	}

	return nil

}
func parsePasswordFromEnv() error {
	viper.SetEnvPrefix("password")
	return viper.BindEnv("salt")
}
func parseJWTFromEnv() error {
	viper.SetEnvPrefix("jwt")
	return viper.BindEnv("signingkey")
}
