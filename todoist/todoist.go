package todoist

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	endpoint = "https://api.todoist.com/rest/v1/"
)

type CreateTaskBody struct {
	Content string `json:"content"`
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
