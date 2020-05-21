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

	log "github.com/sirupsen/logrus"

	"github.com/arunvm/mind/config"
	"github.com/arunvm/mind/todoist"
	"github.com/spf13/cobra"
)

var task *string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds task to todoist inbox",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(*task) == 0 {
			return errors.New("Please provide task description")
		}

		cfg, err := config.ReadConfigFile()
		if err != nil {
			log.Printf("Error when reading config file \n %v", err)
			return err
		}

		todoistClient := todoist.NewClient(cfg.TodoistToken)

		err = todoistClient.CreateTask(*task)
		if err != nil {
			log.Printf("Error when creating task\n%v", err)
			return err
		}

		return nil
	},
}

func init() {
	todoistCmd.AddCommand(addCmd)

	task = addCmd.Flags().String("task", "", "Provide content for task")
}
