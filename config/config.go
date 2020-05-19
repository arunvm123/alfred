package config

import (
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	OutputFormat string `json:"output_format"`
}

//Read the configuration
func ReadConfigFile() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		log.Println("Error when retrieving config path\n %v", err)
		return nil, err
	}

	_, err = os.Stat(configPath)
	if err != nil {
		log.Printf("Error when checking if file exists\n%v", err)
		return nil, err
	}

	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		OutputFormat: viper.GetString("output_format"),
	}, nil
}

//Write data of config to the configuration file
func CreateConfigFile(cfg *Config) error {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Printf("Error when fetching home directory\n%v", err)
		return err
	}

	_, err = os.Create(path.Join(home, ".mind.yaml"))
	if err != nil {
		return err
	}

	viper.SetConfigName(".mind")
	viper.AddConfigPath(home)
	viper.Set("output_format", cfg.OutputFormat)
	return viper.WriteConfig()
}

func getConfigPath() (string, error) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Printf("Error when fetching home directory\n%v", err)
		return "", err
	}

	return path.Join(home, ".mind.yaml"), nil

}
