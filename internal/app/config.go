package app

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type SignerConfig struct {
	SigningKey string
}

type AuthConfig struct {
	AccessTokenTTL time.Duration `yaml:"accessTokenTTL" env-default:"24h"`
}

type RepoConfig struct {
	Type int `yaml:"type"`
}

type HTTPServerConfig struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Config struct {
	AuthConfig       `yaml:"AuthConfig"`
	RepoConfig       `yaml:"RepoConfig"`
	SignerConfig     `yaml:"-"`
	HTTPServerConfig `yaml:"HTTPServerConfig"`
}

func MustLoad() *Config {
	//configPath := flag.String("conf", "config/configs.yml", "Путь до файла конфигурации. Чтобы использовать переменную окружения - необходимо передать ее название при запуске программы: -conf=$CONFIG_PATH")
	//flag.Parse()
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

	return &cfg
}

func PrintConf(conf *Config) {
	logrus.Printf("%+x\n%+x\n%+x\n%+x\n", conf.SignerConfig, conf.RepoConfig, conf.AuthConfig, conf.HTTPServerConfig)
}
