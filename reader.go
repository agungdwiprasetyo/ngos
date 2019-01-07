package ngos

import (
	"encoding/csv"
	"io"
)

// Read function
func Read(reader csv.Reader) ([]string, error) {
	var lines []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		lines = append(lines, line...)

	}

	return lines, nil
}
