package parser

import "time"

// A representation of a log message (Might add-on or change the structure later)

type LogMessage struct {
	TimeStamp time.Time
	Level     string
	Message   string
}
