package types

import (
	"encoding/json"
	"os"
	"time"

	logger "Telegram/Core/Log"
)

func WaitForChanges(c *Config, path string) {
	for {
		LoadJson(path, c)

		time.Sleep(time.Millisecond * 100)
	}
}

func LoadJson(file string, c any) {
	jsonFile, err := os.Open(file)

	if err != nil {
		logger.LogPanic(logger.GetStackTrace(), "%v", err)
	}

	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&c)

	if err != nil {
		logger.LogPanic(logger.GetStackTrace(), "%v", err)
	}
}
