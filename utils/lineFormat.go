package utils

import (
	"Finder/globals"
	"fmt"
	"time"
)

func GetHeaderLine(path string) string {
	return fmt.Sprintf("RequestDate: %s,\tSearchFilter: %s,\tcurrentPath: %s",
		time.Now().Format("02.01.2006 15:04:05"),
		globals.GetFilter(),
		path)
}

func GetResultLine(path string, line string, content string) string {
	return fmt.Sprintf("%d - FilePath: %s,\tLine: %s,\tContent: %s", globals.IncrementLineCounter(), path, line, content)
}

func GetErrorLine(err error) string {
	return fmt.Sprintf("%d - Error is happened: %s\t", globals.IncrementLineCounter(), err.Error())
}
func GetEndLine() string {
	if globals.GetLineCounter() > 0 {
		return fmt.Sprintf("End, total %d lines", globals.GetLineCounter())
	}
	return "Nothing found"
}
