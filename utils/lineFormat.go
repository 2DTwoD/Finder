package utils

import (
	"fmt"
	"time"
)

func GetHeaderLine(filter string) string {
	return fmt.Sprintf("RequestDate: %s,\tSearchFilter: %s", time.Now().Format("02.01.2006 15:04:05"),
		filter)
}

func GetResultLine(path string, line string, content string) string {
	return fmt.Sprintf("FilePath: %s,\tLine: %s,\tContent: %s", path, line, content)
}

func GetErrorLine(err error) string {
	return "Error is happened:\t" + err.Error()
}
