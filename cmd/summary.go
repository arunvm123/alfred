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
	"fmt"

	log "github.com/sirupsen/logrus"

	"os"
	"time"

	"github.com/arunvm/alfred/summary"
	"github.com/spf13/cobra"
)

var date *string

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Retrieves summary of alfred command usage for the specified date",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := time.Parse("2006-01-02", *date)
		if err != nil {
			log.Printf("Error when parsing date\n %v", err)
			return errors.New("Please provide date in 'yyyy-mm-dd' format")
		}

		summaryData, err := summary.GetData(*date)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("No session summary for given date")
				return nil
			}
			log.Printf("Error when retrieving session summary\n%v", err)
			return err
		}

		for _, s := range *summaryData {
			fmt.Println(s)
		}

		return nil
	},
	// PostRunE: func(cmd *cobra.Command, args []string) error {
	// 	err := summary.Save("session", "summary", fmt.Sprintf(summary.SessionSummaryFormat, *date))
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil
	// },
}

func init() {
	sessionCmd.AddCommand(summaryCmd)

	date = summaryCmd.Flags().String("date", "", "Provide date in the format yyyy-mm-dd to retrieve the specified date's summary")
}
