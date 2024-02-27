package main

import (
	"Finder/globals"
	"Finder/pathEntry"
	"Finder/utils"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	globals.SetCurrentFileName(filepath.Base(os.Args[0]))
	globals.SetFilter("wtf") //strings.TrimSuffix(globals.GetCurrentFileName(), filepath.Ext(globals.GetCurrentFileName()))
	globals.SetResultFileName(utils.GetResultFileName())
	resultFile, err := os.Create(globals.GetResultFileName())
	if err != nil {
		log.Fatal(err)
	}
	defer func(resultFile *os.File) {
		err := resultFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resultFile)

	utils.WriteLine(resultFile, utils.GetHeaderLine(utils.GetAbsolutePath("./")))

	var dirEntriesWithPath = make([]*pathEntry.DirEntryWithPath, 0)

	dirEntries, err := os.ReadDir("D:/Virtual Machines/") //"./"
	if err != nil {
		log.Fatal(err)
	}
	for _, dirEntry := range dirEntries {
		dirEntriesWithPath = append(dirEntriesWithPath, pathEntry.New(&dirEntry, "D:/Virtual Machines/")) //"./"
	}
	waitChan := make(chan bool, runtime.NumCPU())
	var i = 0

	for {
		waitChan <- true
		if i+1 > len(dirEntriesWithPath) {
			<-waitChan
			break
		}
		dirEntryWithPath := dirEntriesWithPath[i]
		i++
		go func() {
			if dirEntryWithPath.IsDir() {
				globals.GetMutex().Lock()
				if strings.Contains(dirEntryWithPath.Name(), globals.GetFilter()) {
					utils.WriteLine(
						resultFile,
						utils.GetResultLine(
							utils.GetAbsolutePath(dirEntryWithPath.PathWithName()),
							"Folder name",
							dirEntryWithPath.Name()))
				}

				dirEntryWithPath.AppendPath()
				moreDirEntries, _ := os.ReadDir(dirEntryWithPath.Path())
				for _, dirEntry := range moreDirEntries {
					dirEntriesWithPath = append(dirEntriesWithPath, pathEntry.New(&dirEntry, dirEntryWithPath.Path()))
				}
				globals.GetMutex().Unlock()
				<-waitChan
				return
			}
			utils.SearchScript(dirEntryWithPath, resultFile, waitChan)
		}()
	}
	for len(waitChan) != 0 {
	}
	utils.WriteLine(resultFile, utils.GetEndLine())
}
