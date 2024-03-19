package parser

import "strings"

func ParseLogFile(filePath string) ([]*LogMessage, error) {
	file, err := OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)

	return nil, err
}

// FileType determines the format of the file based on its extension
func FileType(filePath string) string {
	// Get the file extension
	parts := strings.Split(filePath, ".")
	if len(parts) < 2 {
		// No file extension found
		return "unknown"
	}
	extension := parts[len(parts)-1]

	// Determine the file format based on the extension
	switch extension {
	case "json":
		return "json"
	case "csv":
		return "csv"
	default:
		return "text"
	}
}
