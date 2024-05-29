package config

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port        string `yaml:"port"`
		Host        string `yaml:"host"`
		IdleTimeout int    `yaml:"idleTimeout"`
		RWTimout    int    `yaml:"rwTimout"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"dbname"`
		Port     uint16 `yaml:"port"`
		Host     string `yaml:"host"`
		ConnPull int    `yaml:"connectionPull"`
	} `yaml:"database"`
	Log struct {
		Level string `yaml:"level"`
	}
	Logger *slog.Logger
}

func ReadConfig() *Config {
	cfg := Config{}
	if err := cfg.readFile("./config/config.yaml"); err != nil {
		os.Exit(2)
	}
	return &cfg
}

func (src *Config) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		src.Database.Host,
		src.Database.Port,
		src.Database.Username,
		src.Database.Password,
		src.Database.DB,
	)
}

func (src *Config) readFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return src.LogError("readFile:", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(src)
	if err != nil {
		return src.LogError("decodeYaml:", err)
	}
	return nil
}

func (src *Config) LogLevel() slog.Level {
	switch src.Log.Level {
	case "debug":
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}
