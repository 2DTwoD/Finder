package utils

import (
	"Finder/globals"
	"Finder/pathEntry"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const charsAround int = 20

func SearchScript(dirEntryWithPath *pathEntry.DirEntryWithPath, resultFile *os.File, waitChan chan bool) {
	if dirEntryWithPath.Name() == globals.GetResultFileName() || dirEntryWithPath.Name() == globals.GetCurrentFileName() {
		<-waitChan
		return
	}
	matches := make([]string, 0)
	if dirEntryWithPath.IsDir() {
		if strings.Contains(strings.ToLower(dirEntryWithPath.Name()), globals.GetFilter()) {
			matches = append(matches,
				GetResultLine(GetAbsolutePath(dirEntryWithPath.Path()), "Folder name", dirEntryWithPath.Name()))
		}
	} else {
		matches = searchInFile(dirEntryWithPath.PathWithName(), dirEntryWithPath.Name())
	}
	for _, item := range matches {
		globals.GetMutex().Lock()
		WriteLine(resultFile, fmt.Sprintf("%d - %s", globals.IncrementLineCounter(), item))
		globals.GetMutex().Unlock()
	}
	<-waitChan
}

func searchInFile(path string, name string) []string {
	result := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		result = append(result, GetErrorLine(err))
		return result
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			result = append(result, GetErrorLine(err))
		}
	}(file)

	abs := GetAbsolutePath(path)
	if strings.Contains(strings.ToLower(name), globals.GetFilter()) {
		result = append(result, GetResultLine(abs, "File name", name))
	}
	scanner := bufio.NewScanner(file)
	line := 1
	for scanner.Scan() {
		carriage := 0
		for {
			currentSearchLine := scanner.Text()[carriage:]
			if carriage >= len(scanner.Text()) {
				break
			}

			if len(scanner.Text()[carriage:]) != len(strings.ToLower(scanner.Text()[carriage:])) {
				println(scanner.Text()[carriage:])
				println(strings.ToLower(scanner.Text()[carriage:]))
			}

			if strings.Contains(currentSearchLine, globals.GetFilter()) {
				index := strings.Index(currentSearchLine, globals.GetFilter())
				start := max(0, carriage+index-charsAround)
				finish := min(carriage+index+len(globals.GetFilter())+charsAround, len(scanner.Text()))
				result = append(
					result,
					GetResultLine(abs, strconv.Itoa(line), fmt.Sprintf("...%s...", scanner.Text()[start:finish])))
				carriage += index + len(globals.GetFilter())
			} else {
				break
			}
		}
		line++
	}
	return result
}
