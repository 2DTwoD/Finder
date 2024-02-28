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

const startPath string = "./"
const startFilter string = "wtf"

func main() {
	log.Println("Start")

	globals.SetCurrentFileName(filepath.Base(os.Args[0]))
	if startFilter == "" {
		globals.SetFilter(strings.ToLower(strings.TrimSuffix(globals.GetCurrentFileName(), filepath.Ext(globals.GetCurrentFileName()))))
	} else {
		globals.SetFilter(strings.ToLower(startFilter))
	}
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

	utils.WriteLine(resultFile, utils.GetHeaderLine(utils.GetAbsolutePath(startPath)))

	var dirEntriesWithPath = make([]*pathEntry.DirEntryWithPath, 0)

	dirEntries, err := os.ReadDir(startPath)
	if err != nil {
		log.Fatal(err)
	}
	dirEntriesWithPath = getEntryPathSlice(dirEntries, startPath)

	var i int

	log.Println("Creating folders tree...")

	for {
		if i >= len(dirEntriesWithPath) {
			break
		}
		dirEntryWithPath := dirEntriesWithPath[i]
		i++
		if dirEntryWithPath.IsDir() {
			dirEntryWithPath.AppendPath()
			moreDirEntries, _ := os.ReadDir(dirEntryWithPath.Path())
			dirEntriesWithPath = append(dirEntriesWithPath, getEntryPathSlice(moreDirEntries, dirEntryWithPath.Path())...)
		}
	}

	log.Println("Search for matches...")

	waitChan := make(chan bool, runtime.NumCPU())
	for _, dirEntryWithPath := range dirEntriesWithPath {
		waitChan <- true
		go utils.SearchScript(dirEntryWithPath, resultFile, waitChan)
	}

	for len(waitChan) != 0 {
	}

	log.Println("Finish")

	utils.WriteLine(resultFile, utils.GetEndLine())
}

func getEntryPathSlice(dirEntrySlice []os.DirEntry, path string) []*pathEntry.DirEntryWithPath {
	result := make([]*pathEntry.DirEntryWithPath, 0)
	for _, dirEntry := range dirEntrySlice {
		result = append(result, pathEntry.New(&dirEntry, path))
	}
	return result
}
