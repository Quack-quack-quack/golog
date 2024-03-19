package parser

import (
	"errors"
	"strings"
	"time"
)

func ParseLine(line string) (*LogMessage, error) {
	// Seperating the line with the delimiter (currently just a space)
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return nil, errors.New("Invalid log foramt")
	}
	timestamp, err := ParseTimeStamp(parts[0])
	if err != nil {
		return nil, err
	}

	return &LogMessage{
		TimeStamp: timestamp,
		Level:     parts[1],
		Message:   parts[2],
	}, nil

}

func ParseTimeStamp(timeStampStr string) (time.Time, error) {
	layout := time.RFC3339

	timeStamp, err := time.Parse(layout, timeStampStr)
	if err != nil {
		return time.Time{}, err
	}
	return timeStamp, nil
}
