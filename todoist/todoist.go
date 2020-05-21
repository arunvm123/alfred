package todoist

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	endpoint = "https://api.todoist.com/rest/v1/"
)

type CreateTaskBody struct {
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

func CreateTask(token, content string) error {
	data := CreateTaskBody{
		Content: content,
	}

	byteData, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	url := endpoint + "tasks"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func GetTasks(token, date string) (*[]Task, error) {
	url := endpoint + "tasks"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tasks []Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
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
