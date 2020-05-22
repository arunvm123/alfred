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
	"time"

	"github.com/arunvm/mind/config"
	"github.com/arunvm/mind/summary"
	"github.com/arunvm/mind/todoist"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfigFile()
		if err != nil {
			log.Printf("Error when reading config file \n %v", err)
			return err
		}

		todoistClient := todoist.NewClient(cfg.TodoistToken)

		tasks, err := todoistClient.GetTasksForDate(time.Now().Format("2006-01-02"))
		if err != nil {
			log.Printf("Error when fetching tasks \n %v", err)
			return err
		}

		if len(*tasks) == 0 {
			fmt.Println("No tasks due for today")
		}

		for _, task := range *tasks {
			fmt.Println(task)
		}

		return nil
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		err := summary.Save("todoist", "today", "")
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	todoistCmd.AddCommand(todayCmd)
}
