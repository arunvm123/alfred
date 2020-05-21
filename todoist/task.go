package todoist

import (
	"bytes"
	"encoding/json"
	"time"
)

type createTaskBody struct {
	Content string `json:"content"`
}

type Task struct {
	Content string `json:"content"`
	Due     *Due   `json:"due"`
}

type Due struct {
	Date   string `json:"date"`
	String string `json:"string"`
}

func (c *Client) CreateTask(task string) error {
	data := createTaskBody{
		Content: task,
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
