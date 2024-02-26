package utils

import (
	"errors"
	"fmt"
	"os"
)

func GetResultFileName() string {
	var newName string
	count := 0
	for {
		newName = fmt.Sprintf("result%d.txt", count)
		_, err := os.Stat(newName)
		if err == nil {
			count++
			continue
		}
		if errors.Is(err, os.ErrNotExist) {
			return newName
		}
	}
}
