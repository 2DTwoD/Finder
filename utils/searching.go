package utils

import (
	"Finder/classDirEntryWithPath"
	"Finder/structures"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const charsAround int = 20

func SearchScript(dirEntryWithPath *classDirEntryWithPath.DirEntryWithPath, resultFile *os.File, waitChan chan bool,
	allNames *structures.Names, mutex *sync.Mutex) {
	if dirEntryWithPath.Name() == allNames.Result || dirEntryWithPath.Name() == allNames.Current {
		<-waitChan
		return
	}
	matches := searchInFile(dirEntryWithPath.PathWithName(), dirEntryWithPath.Name(), allNames.Filter)
	for _, item := range matches {
		mutex.Lock()
		resultFile.WriteString(item + "\n")
		mutex.Unlock()
	}
	<-waitChan
}

func searchInFile(path string, name string, filter string) []string {
	result := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		result = append(result, "Error is happened: "+err.Error())
		return result
	}
	defer file.Close()
	abs, err := filepath.Abs(file.Name())
	if strings.Contains(name, filter) {
		result = append(result,
			fmt.Sprintf("FilePath: %s, Line: In name, Content: %s",
				abs, name))
	}
	scanner := bufio.NewScanner(file)
	line := 1
	for scanner.Scan() {
		carret := 0
		for {
			currentSearchLine := scanner.Text()[carret:]
			if carret > len(scanner.Text()) {
				break
			}
			if strings.Contains(currentSearchLine, filter) {
				index := strings.Index(currentSearchLine, filter)
				start := max(0, carret+index-charsAround)
				finish := min(carret+index+len(filter)+charsAround, len(scanner.Text()))
				if err != nil {
					abs = "ABS_PATH_ERROR"
				}
				result = append(result,
					fmt.Sprintf("FilePath: %s, Line: %d, Content: %s",
						abs, line, fmt.Sprintf("...%s...", scanner.Text()[start:finish])))
				carret += index + len(filter)
			} else {
				break
			}
		}
		line++
	}
	return result
}
