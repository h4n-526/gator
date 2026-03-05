package config

import (
  "encoding/json"
  "os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
  DBURL           string `json:"db_url"`
  CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
  cfg.CurrentUserName = username
  return write(*cfg)
}

func (cfg *Config) SetDBURL(dbURL string) error {
  cfg.DBURL = dbURL
  return write(*cfg)
}

func Read() (Config, error) {
  path, err := getConfigFilePath()
  if err != nil {
    return Config{}, err
  }
  file, err := os.Open(path)
  if err != nil {
    return Config{}, err
  }
  defer file.Close()
  decoder := json.NewDecoder(file)
  var cfg Config
  err = decoder.Decode(&cfg)
  return cfg, nil
}

func getConfigFilePath() (string, error) {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    return "", err
  }
  return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
  path, err := getConfigFilePath()
  if err != nil {
    return err
  }
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()
  return json.NewEncoder(file).Encode(cfg)
}
