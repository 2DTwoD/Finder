package globals

import (
	"sync"
)

type Names struct {
	current string
	result  string
	filter  string
}

var allNames Names
var mutex sync.Mutex
var lineCounter int

func GetFilter() string {
	return allNames.filter
}

func SetFilter(value string) {
	allNames.filter = value
}

func GetCurrentFileName() string {
	return allNames.current
}

func SetCurrentFileName(value string) {
	allNames.current = value
}

func GetResultFileName() string {
	return allNames.result
}

func GetMutex() *sync.Mutex {
	return &mutex
}

func SetResultFileName(value string) {
	allNames.result = value
}

func IncrementLineCounter() int {
	lineCounter++
	return lineCounter
}

func GetLineCounter() int {
	return lineCounter
}
