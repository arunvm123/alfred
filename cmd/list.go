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
	"encoding/json"
	"fmt"
	"os"

	"github.com/arunvm/alfred/config"
	"github.com/arunvm/alfred/summary"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieves list of all channels and their channel ID. Channel ID is required to send messages to that channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfigFile()
		if err != nil {
			log.Printf("Error when reading config file\n%v", err)
			return err
		}

		slackClient := slack.New(cfg.SlackToken)
		conversations, _, err := slackClient.GetConversations(&slack.GetConversationsParameters{
			Types:           []string{"im", "public_channel", "private_channel"},
			ExcludeArchived: "true",
		})
		if err != nil {
			log.Printf("Error when fetching channels \n", err)
			return err
		}

		channelInfo, err := getChannelInfo(slackClient, &conversations)
		if err != nil {
			log.Printf("Error when formatting channel info \n", err)
			return err
		}

		for _, channel := range *channelInfo {
			fmt.Println(channel)
		}

		return nil
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		err := summary.Save("slack", "list", "")
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	slackCmd.AddCommand(listCmd)
}

type channelInfo struct {
	ChannelID string `json:"channel_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

func getChannelInfo(slackClient *slack.Client, conversations *[]slack.Channel) (*[]channelInfo, error) {
	var ci []channelInfo

	// This is a helper variable to map userID to IM details in the channel info array.
	// Used to store the name after fetching all the user info in one call, rather than call for each user
	userToChannel := make(map[string]int)
	var userIDs []string

	for _, channel := range *conversations {
		chanInfo := channelInfo{
			ChannelID: channel.ID,
		}

		if channel.IsIM {
			userIDs = append(userIDs, channel.User)
			chanInfo.Type = "IM"

			userToChannel[channel.User] = len(ci)
		} else if channel.IsChannel {
			chanInfo.Name = channel.Name
			chanInfo.Type = "Channel"
		}

		ci = append(ci, chanInfo)
	}

	users, err := slackClient.GetUsersInfo(userIDs...)
	if err != nil {
		log.Printf("Error when fetching info of all users\n%v", err)
		return nil, err
	}

	for _, user := range *users {
		ci[userToChannel[user.ID]].Name = user.Name
	}

	return &ci, nil
}

func (c channelInfo) String() string {
	cfg, err := config.ReadConfigFile()
	if err != nil {
		log.Printf("Error when reading config file%v", err)
		os.Exit(1)
	}

	if cfg.OutputFormat == "plain text" {
		return fmt.Sprintf("ID: %s, Name: %s, Type: %s", c.ChannelID, c.Name, c.Type)
	} else if cfg.OutputFormat == "json" {
		// Convert structs to JSON.
		data, err := json.MarshalIndent(c, "", "\t")
		if err != nil {
			log.Printf("Error when marshalling channel data\n %v", err)
			os.Exit(1)
		}

		return fmt.Sprintf("%s", string(data))
	}

	return ""

}
