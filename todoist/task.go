package todoist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/arunvm/mind/config"
)

type createTaskBody struct {
	Content string `json:"content"`
	DueDate string `json:"due_date"`
}

type Task struct {
	Content string `json:"content"`
	Due     *Due   `json:"due"`
}

type Due struct {
	Date   string `json:"date"`
	String string `json:"string"`
}

func (c *Client) CreateTask(task, date string) error {
	data := createTaskBody{
		Content: task,
		DueDate: date,
	}

	byteData, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	err = c.postMethod("/tasks", bytes.NewReader(byteData), nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetTasksForDate(date string) (*[]Task, error) {
	var tasks []Task
	err := c.getMethod("/tasks", &tasks)
	if err != nil {
		return nil, err
	}

	filteredTasks, err := filterTasksByDate(&tasks, date)
	if err != nil {
		return nil, err
	}

	return filteredTasks, nil
}

func filterTasksByDate(tasks *[]Task, date string) (*[]Task, error) {
	var filteredTasks []Task

	givenDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	for _, task := range *tasks {
		if task.Due != nil {
			taskDate, err := time.Parse("2006-01-02", task.Due.Date)
			if err != nil {
				return nil, err
			}

			if taskDate.Equal(givenDate) {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	return &filteredTasks, nil
}

func (t Task) String() string {
	cfg, err := config.ReadConfigFile()
	if err != nil {
		log.Printf("Error when reading config file\n%v", err)
		os.Exit(1)
	}

	if cfg.OutputFormat == "plain text" {
		return fmt.Sprintf("Task: %v", t.Content)
	} else if cfg.OutputFormat == "json" {
		// Convert structs to JSON.
		data, err := json.MarshalIndent(t, "", "\t")
		if err != nil {
			log.Printf("Error when marshalling task data\n %v", err)
			os.Exit(1)
		}

		return fmt.Sprintf("%s", string(data))
	}

	return ""
}
