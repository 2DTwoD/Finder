package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

const charsAround int = 20

type Names struct {
	current string
	result  string
	filter  string
}

var names Names
var matches = make([]string, 0)
var mutex sync.Mutex

func main() {
	names.current = filepath.Base(os.Args[0])
	names.filter = "111" //strings.TrimSuffix(names.current, filepath.Ext(names.current))
	names.result = getFileName()
	resultFile, err := os.Create(names.result)
	if err != nil {
		panic("Unable to create file")
		os.Exit(1)
	}
	defer resultFile.Close()
	resultFile.WriteString(
		fmt.Sprintf("RequestDate: %s, SearchFilter: %s\n", time.Now().Format("02.01.2006 15:04:05"),
			names.filter))

	filesInfo, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	waitChan := make(chan bool, runtime.NumCPU())
	for _, fileInfo := range filesInfo {
		waitChan <- true
		go searchScript(fileInfo, resultFile, waitChan)
	}
	for len(waitChan) != 0 {
	}
}

func searchScript(fileInfo fs.FileInfo, resultFile *os.File, waitChan chan bool) {
	if fileInfo.IsDir() || fileInfo.Name() == names.result || fileInfo.Name() == names.current {
		<-waitChan
		return
	}
	matches = searchInFile(fileInfo.Name(), names.filter)
	for _, item := range matches {
		mutex.Lock()
		resultFile.WriteString(item + "\n")
		mutex.Unlock()
	}
	<-waitChan
}

func searchInFile(name string, filter string) []string {
	result := make([]string, 0)
	file, err := os.Open(name)
	if err != nil {
		result[0] = "Error is happened: " + err.Error()
		return result
	}
	defer file.Close()
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
				finish := min(carret+index+len(names.filter)+charsAround, len(scanner.Text()))
				abs, err := filepath.Abs(file.Name())
				if err != nil {
					abs = "ABS_PATH_ERROR"
				}
				result = append(result,
					fmt.Sprintf("FilePath: %s, Line: %d, Content: %s",
						abs, line, fmt.Sprintf("...%s...", scanner.Text()[start:finish])))
				carret += index + len(names.filter)
			} else {
				break
			}
		}
		line++
	}
	return result
}

func getFileName() string {
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
