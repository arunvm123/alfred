# alfred

  Alfred is a command line tool that lets you interact with slack and todoist. You will be able to list all the channels in slack and send messages to them. The available commands for todoist will help you to add tasks for any specified day and view all tasks for today. Alfred also provides commands to configure output format.

  Alfred is built with the very popular library Cobra, which is also used in other Go projects such as Kubernetes, Hugo, and Github CLI. The configuration management is setup with Viper.

## Install the CLI

  Move to the root of the directory and run the following command

  ``` 
  go install 
  ```

## Basic Usage

  ```

  A handy tool to carry out your day to day workUsage:
  alfred [flags]
  
  alfred [command]
  Available Commands:
  configure   Configure the ouput format. It can be either 'json' or 'plain text' 
  
  help        Help about any command
  session     Top level command for session info
  slack       Top level command for slack with flag to authorise user to slack
  todoist     Top level command for slack with flag to authorise user to todoist
  Flags:
  -h, --help     help for alfred
  
  -t, --toggle   Help message for toggle
  Use "alfred [command] --help" for more information about a command.
  ```
  