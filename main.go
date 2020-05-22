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
package main

import (
	"os"
	"path"

	"github.com/arunvm/mind/cmd"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Error when fetching home directory\n%v", err)
	}

	file, err := os.OpenFile(path.Join(home, ".mind.error.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error when opening/creating error log file \n %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.SetFormatter(&log.JSONFormatter{})

	cmd.Execute()
}
