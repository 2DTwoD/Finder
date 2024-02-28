package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetResultFileName() string {
	var newName string
	count := 0
	for {
		newName = fmt.Sprintf("!result%d.txt", count)
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

func WriteLine(file *os.File, line string) {
	_, err := file.WriteString(fmt.Sprintf("%s\n", line))
	if err != nil {
		log.Fatal(err)
	}
}

func GetAbsolutePath(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
	}
	return abs
}
