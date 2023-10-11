package utils

import (
	"bufio"
	"encoding/base64"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"

	logger "Telegram/Core/Log"
)

var (
	mutex sync.Mutex
)

// Appends line to file
func AppendLine(filepath string, s string, m *sync.Mutex) error {
	m.Lock()
	defer m.Unlock()
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.WriteString(s + "\n"); err != nil {
		return err
	}

	return nil
}

// Read lines of file
func Readlines(path string) []string {
	var Lines []string

	readFile, err := os.Open(path)

	if err != nil {
		logger.LogPanic(logger.GetStackTrace(), "%v", err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		Lines = append(Lines, fileScanner.Text())
	}

	return Lines
}

// Delete files contents
func DestroyFile(filepath string) {
	mutex.Lock()
	defer mutex.Unlock()
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)

	if err != nil {
		logger.LogPanic(logger.GetStackTrace(), "%v", err)
	}

	defer file.Close()
}

// Deletes line from file
func DeleteLine(filepath, s string) {
	mutex.Lock()
	defer mutex.Unlock()

	lines := Readlines(filepath)

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)

	if err != nil {
		logger.LogPanic(logger.GetStackTrace(), "%v", err)
	}

	for _, line := range lines {
		if !strings.Contains(line, s) {
			file.WriteString(line + "\n")
		}
	}
}

// Checks if an array contains "s", does not check for equality; checks for similarities: strings.Contains().
func ArrayContains(arr []string, s string) bool {
	for _, v := range arr {
		if strings.Contains(v, s) {
			return true
		}
	}

	return false
}

// Removes duplicates from an array
func RemoveDupes(s []string) (rl []string) {
	contains := func(s string, sl []string) bool {
		for _, v := range sl {
			if v == s {
				return true
			}
		}

		return false
	}

	for _, v := range s {
		if !contains(v, rl) {
			rl = append(rl, v)
		}
	}

	return
}

// Removes item from an array. Pointer only.
func RemoveFromArray(arrP *[]string, s string) {
	arr := *arrP

	for i, v := range arr {
		if strings.Contains(v, s) {
			*arrP = append(arr[:i], arr[i+1:]...)
			return
		}
	}
}

func GetAllFiles(path string) []os.DirEntry {
	entries, err := os.ReadDir(path)

	if err != nil {
		logger.LogPanic(logger.GetStackTrace(), "Could not get all the files: %v, Error: %v", path, err.Error())
	}

	return entries
}

func ShuffleArray(array *[]string) {
	slice := *array

	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ImageToB64(path string) string {
	file, err := os.Open(path)

	if err != nil {
		logger.LogPanic(logger.GetStackTrace(), err.Error())
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		logger.LogPanic(logger.GetStackTrace(), err.Error())
	}

	return "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(data)
}
