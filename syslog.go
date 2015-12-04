package event

import "time"

type SyslogV1 struct {
	Client    string    `json:"client"`
	Content   string    `json:"content"`
	Facility  int       `json:"facility"`
	Hostname  string    `json:"hostname"`
	Priority  int       `json:"priority"`
	Severity  int       `json:"severity"`
	Tag       string    `json:"tag"`
	Timestamp time.Time `json:"timestamp"`
}
