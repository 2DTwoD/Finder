package main

import (
	"Finder/classDirEntryWithPath"
	"Finder/structures"
	"Finder/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var allNames structures.Names
var mutex sync.Mutex

func main() {
	allNames.Current = filepath.Base(os.Args[0])
	allNames.Filter = "111" //strings.TrimSuffix(names.current, filepath.Ext(names.current))
	allNames.Result = utils.GetResultFileName()
	resultFile, err := os.Create(allNames.Result)
	if err != nil {
		panic("Unable to create file")
		os.Exit(1)
	}
	defer resultFile.Close()
	resultFile.WriteString(
		fmt.Sprintf("RequestDate: %s, SearchFilter: %s\n", time.Now().Format("02.01.2006 15:04:05"),
			allNames.Filter))

	var dirEntrysWithPath = make([]*classDirEntryWithPath.DirEntryWithPath, 0)

	dirEntrys, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, dirEntry := range dirEntrys {
		dirEntrysWithPath = append(dirEntrysWithPath, classDirEntryWithPath.New(&dirEntry, "./"))
	}
	waitChan := make(chan bool, runtime.NumCPU())
	var i = 0
	for {
		if i+1 > len(dirEntrysWithPath) {
			break
		}
		dirEntryWithPath := dirEntrysWithPath[i]
		i++
		if dirEntryWithPath.IsDir() {
			dirEntryWithPath.AppendPath()
			moreDirEntrys, _ := os.ReadDir(dirEntryWithPath.Path())
			for _, dirEntry := range moreDirEntrys {
				dirEntrysWithPath = append(dirEntrysWithPath, classDirEntryWithPath.New(&dirEntry, dirEntryWithPath.Path()))
			}
			continue
		}
		waitChan <- true
		go utils.SearchScript(dirEntryWithPath, resultFile, waitChan, &allNames, &mutex)
	}
	for len(waitChan) != 0 {
	}
}
