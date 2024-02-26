package utils

import (
	"Finder/pathEntry"
	"Finder/structures"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const charsAround int = 20

func SearchScript(dirEntryWithPath *pathEntry.DirEntryWithPath, resultFile *os.File, waitChan chan bool,
	allNames *structures.Names, mutex *sync.Mutex) {
	if dirEntryWithPath.Name() == allNames.Result || dirEntryWithPath.Name() == allNames.Current {
		<-waitChan
		return
	}
	matches := searchInFile(dirEntryWithPath.PathWithName(), dirEntryWithPath.Name(), allNames.Filter)
	for _, item := range matches {
		mutex.Lock()
		WriteLine(resultFile, item)
		mutex.Unlock()
	}
	<-waitChan
}

func searchInFile(path string, name string, filter string) []string {
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
	if strings.Contains(name, filter) {
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
			if strings.Contains(currentSearchLine, filter) {
				index := strings.Index(currentSearchLine, filter)
				start := max(0, carriage+index-charsAround)
				finish := min(carriage+index+len(filter)+charsAround, len(scanner.Text()))
				result = append(
					result,
					GetResultLine(abs, strconv.Itoa(line), fmt.Sprintf("...%s...", scanner.Text()[start:finish])))
				carriage += index + len(filter)
			} else {
				break
			}
		}
		line++
	}
	return result
}
