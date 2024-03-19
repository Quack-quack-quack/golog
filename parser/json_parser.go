package parser

import "encoding/json"

func ParseJSONLine(line string) (*LogMessage, error) {
	var logMsg LogMessage
	err := json.Unmarshal([]byte(line), &logMsg)
	if err != nil {
		return nil, err
	}
	return &logMsg, nil
}

func ReadJSON() {}
