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
	"errors"

	"github.com/arunvm/alfred/config"
	"github.com/arunvm/alfred/summary"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var todoist_token *string

// todoistCmd represents the todoist command
var todoistCmd = &cobra.Command{
	Use:   "todoist",
	Short: "Top level command for slack with flag to authorise user to todoist",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(*todoist_token) == 0 {
			return errors.New("Provide todoist token")
		}

		cfg, err := config.ReadConfigFile()
		if err != nil {
			log.Printf("Error when reading config file\n%v", err)
			return err
		}

		cfg.TodoistToken = *todoist_token
		err = cfg.SaveConfig()
		if err != nil {
			log.Printf("Error when saving todoist token \n %v", err)
			return err
		}

		return nil
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		err := summary.Save("todoist", "auth", "")
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(todoistCmd)

	todoist_token = todoistCmd.Flags().String("auth_token", "", "Add token to authorise to todoist")
}
