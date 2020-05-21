package summary

import (
	"encoding/csv"
	"os"
	"path"
	"time"

	"github.com/mitchellh/go-homedir"
)

type Summary struct {
	Command    string
	SubCommand string
	Args       string
	Time       string
}

var (
	ConfigureStringFormat = "output_format=%s"
	SlackSendFormat       = "message=%s|channel=%s"
	TodoistAddFormat      = "task=%s"
	SessionSummaryFormat  = "date=%s"
)

func Save(command, subcommand, args string) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	summaryDirectory := path.Join(home, ".mind_summary")

	err = os.MkdirAll(summaryDirectory, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(summaryDirectory, time.Now().Format("2006-01-02")+".csv"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	err = w.Write([]string{command, subcommand, args, time.Now().Format("15:04:05")})
	if err != nil {
		return err
	}

	w.Flush()

	return nil
}

func GetData(date string) (*[]Summary, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path.Join(home, ".mind_summary", date+".csv"), os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	var summaries []Summary
	// Loop through lines & turn into object
	for _, line := range lines {
		s := Summary{
			Command:    line[0],
			SubCommand: line[1],
			Args:       line[2],
			Time:       line[3],
		}
		summaries = append(summaries, s)
	}

	return &summaries, nil
}
