/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mind",
	Short: "A handy tool to carry out your day to day work",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("In root command run")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mind.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig creates a config file if it does not exist
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Printf("Error when fetching home directory\n%v", err)
		os.Exit(1)
	}

	configPath := path.Join(home, ".mind.yaml")

	if _, err := os.Stat(configPath); err == nil {
		return
	}

	_, err = os.Create(configPath)
	if err != nil {
		log.Printf("Error when creating config file\n%v", err)
		os.Exit(1)
	}

	// Search config in home directory with name ".mind".
	viper.AddConfigPath(home)
	viper.SetConfigName(".mind")

	// Setting default value of output format to 'json'
	viper.Set("output_format", "json")
	err = viper.WriteConfig()
	if err != nil {
		log.Printf("Error when writing to config file\n%v", err)
		os.Exit(1)
	}

	return
}
