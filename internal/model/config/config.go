package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	App      string `json:"app"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func NewFromEnv() (*Config, error) {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	return nil, err
	//}

	app := os.Getenv("APP")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")

	return &Config{
		App:      app,
		Port:     port,
		Database: database,
	}, nil
}

func New() (*Config, error) {
	filePath := "./files/secrets.config.json"
	var config Config

	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file contents
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into the Config struct
	err = json.Unmarshal(fileData, &config)
	return &config, err
}

type LocalImage struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

func LoadImages() ([]LocalImage, error) {
	filePath := "./files/images.json"
	var imgs []LocalImage

	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file contents
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into the Config struct
	err = json.Unmarshal(fileData, &imgs)
	return imgs, err
}
