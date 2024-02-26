package main

import (
	"Finder/pathEntry"
	"Finder/structures"
	"Finder/utils"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var allNames structures.Names
var mutex sync.Mutex

func main() {
	allNames.Current = filepath.Base(os.Args[0])
	allNames.Filter = strings.TrimSuffix(allNames.Current, filepath.Ext(allNames.Current))
	allNames.Result = utils.GetResultFileName()
	resultFile, err := os.Create(allNames.Result)
	if err != nil {
		log.Fatal(err)
	}
	defer func(resultFile *os.File) {
		err := resultFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resultFile)

	utils.WriteLine(resultFile, utils.GetHeaderLine(allNames.Filter))

	var dirEntriesWithPath = make([]*pathEntry.DirEntryWithPath, 0)

	dirEntries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, dirEntry := range dirEntries {
		dirEntriesWithPath = append(dirEntriesWithPath, pathEntry.New(&dirEntry, "./"))
	}
	waitChan := make(chan bool, runtime.NumCPU())
	var i = 0
	for {
		if i+1 > len(dirEntriesWithPath) {
			break
		}
		dirEntryWithPath := dirEntriesWithPath[i]
		waitChan <- true
		i++
		go func() {
			if dirEntryWithPath.IsDir() {
				if strings.Contains(dirEntryWithPath.Name(), allNames.Filter) {
					utils.WriteLine(
						resultFile,
						utils.GetResultLine(dirEntryWithPath.PathWithName(), "Folder name", dirEntryWithPath.Name()))
				}

				dirEntryWithPath.AppendPath()
				moreDirEntries, _ := os.ReadDir(dirEntryWithPath.Path())
				for _, dirEntry := range moreDirEntries {
					dirEntriesWithPath = append(dirEntriesWithPath, pathEntry.New(&dirEntry, dirEntryWithPath.Path()))
				}
				<-waitChan
				return
			}
			utils.SearchScript(dirEntryWithPath, resultFile, waitChan, &allNames, &mutex)
		}()
	}
	for len(waitChan) != 0 {
	}
}
