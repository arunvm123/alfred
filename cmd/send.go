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

	"github.com/arunvm/mind/config"
	"github.com/arunvm/mind/summary"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

var channelID *string
var message *string

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends message to the specified channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(*message) == 0 {
			return errors.New("Please provide a message")
		}

		if len(*channelID) == 0 {
			return errors.New("Please provide a channel ID")
		}

		cfg, err := config.ReadConfigFile()
		if err != nil {
			log.Printf("Error when reading config file\n%v", err)
			return err
		}

		slackClient := slack.New(cfg.SlackToken)
		_, _, err = slackClient.PostMessage(*channelID, slack.MsgOptionText(*message, false))
		if err != nil {
			log.Printf("Error when sending message\n%v", err)
			return err
		}

		return nil
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		err := summary.Save("slack", "send", fmt.Sprintf(summary.SlackSendFormat, *message, *channelID))
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	slackCmd.AddCommand(sendCmd)

	channelID = sendCmd.Flags().StringP("channel", "c", "", "Specify channel ID to send message")
	message = sendCmd.Flags().StringP("message", "m", "", "Message to be send to channel")
}
