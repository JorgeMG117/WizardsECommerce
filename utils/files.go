package utils

import (
	"encoding/json"
	"os"
)

func ReadFile(fileName string, data interface{}) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, data)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(fileName string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, file, 0644)
	if err != nil {
		return err
	}
	return nil
}
