package setup

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Spotify struct {
		ClientId string `yaml:"client_id"`
	}
	Musixmatch struct {
		ApiKey string `yaml:"api_key"`
	}
}

var Env = &Config{}

func getConfigFilePath() string {
	configPath := os.Getenv("CLICKARAOKE_CONFIG_FILE")

	if configPath != "" {
		return configPath
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error getting user config dir: %v", err)
	}

	return filepath.Join(configDir, "clickaraoke", "config.yaml")
}

func getConfigFile() ([]byte, error) {
	configPath := getConfigFilePath()

	file, err := os.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func saveConfigFile() {
	configPath := getConfigFilePath()

	_ = os.MkdirAll(filepath.Dir(configPath), 0755)

	file, err := os.Create(configPath)

	if err != nil {
		log.Fatalf("Error creating config file: %v", err)
	}

	defer file.Close()

	bytes, err := yaml.Marshal(Env)

	if err != nil {
		log.Fatalf("Error marshalling config: %v", err)
	}

	_, err = file.Write(bytes)

	if err != nil {
		log.Fatalf("Error writing config file: %v", err)
	}
}

func manualSetup() {
	fmt.Println("Is this your first time running Clickaraoke, let's setup your environment")
	//TODO: Criar o readme :p
	fmt.Println("Any questions, please see the README.md on github ")

	fmt.Print("Please enter your Spotify Client ID: ")
	fmt.Scanln(&Env.Spotify.ClientId)

	fmt.Print("\nPlease enter your Musixmatch API Key: ")
	fmt.Scanln(&Env.Musixmatch.ApiKey)

	fmt.Printf("Config file will be saved at %s\n", getConfigFilePath())

	saveConfigFile()
}

func loadEnv() {
	file, err := getConfigFile()

	if err != nil && errors.Is(err, os.ErrNotExist) {
		manualSetup()
		return
	} else if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	yaml.Unmarshal(file, Env)
}

func Init() {
	loadEnv()
}
