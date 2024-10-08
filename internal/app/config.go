package app

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type MongoConfig struct {
	Username       string `yaml:"username"`
	Password       string `yaml:"-"`
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	DBName         string `yaml:"db_name"`
	CollectionName string `yaml:"collection_name"`
}

type MySQLConfig struct {
	Username    string `yaml:"username"`
	Password    string `yaml:"-"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	DBName      string `yaml:"db_name"`
	MaxOpenConn int    `yaml:"max_open_conn"`
}

type SignerConfig struct {
	SigningKey string
}

type AuthConfig struct {
	AccessTokenTTL time.Duration `yaml:"accessTokenTTL" env-default:"24h"`
	SessionTTL     time.Duration `yaml:"sessionTTL" env-default:"43200s"`
}

type RepoConfig struct {
	Type string `yaml:"type"`
}

type RedisConfig struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

type HTTPServerConfig struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Config struct {
	MySQLConfig      `yaml:"MySQLConfig"`
	MongoConfig      `yaml:"MongoConfig"`
	AuthConfig       `yaml:"AuthConfig"`
	RepoConfig       `yaml:"RepoConfig"`
	SignerConfig     `yaml:"-"`
	HTTPServerConfig `yaml:"HTTPServerConfig"`
	RedisConfig      `yaml:"RedisConfig"`
}

func MustLoad() *Config {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalln(err)
	}

	configPath := os.Getenv("CONFIG_PATH")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logrus.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		logrus.Fatalln(err)
	}

	cfg.SignerConfig.SigningKey = os.Getenv("SIGNING_KEY")
	cfg.MongoConfig.Password = os.Getenv("MONGODB_PASSWORD")
	cfg.MySQLConfig.Password = os.Getenv("MYSQL_PASSWORD")

	return &cfg
}

func PrintConf(conf *Config) {
	logrus.Printf("%+x\n%+x\n%+x\n%+x\n", conf.SignerConfig, conf.RepoConfig, conf.AuthConfig, conf.HTTPServerConfig)
}
