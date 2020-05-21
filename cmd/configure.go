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

	"github.com/arunvm/mind/config"
	"github.com/arunvm/mind/summary"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var output_format *string

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure the ouput format. It can be either 'json' or 'plain text'",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfigFile()
		if err != nil {
			log.Printf("Error when reading config file\n%v", err)
			return err
		}

		cfg.OutputFormat = *output_format
		err = cfg.SaveConfig()
		if err != nil {
			log.Printf("Error when saving output format \n %v", err)
			return err
		}

		return nil
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		err := summary.Save("configure", "", fmt.Sprintf(summary.ConfigureStringFormat, *output_format))
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	output_format = configureCmd.Flags().String("output_format", "", "output format can either be 'json' or 'plain text")
}
