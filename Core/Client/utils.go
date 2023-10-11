package client

import (
	types "Telegram/Core/Types"
	"encoding/json"
	"os"
)

func GetAllSessions(path string) (*types.Session, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var Response *types.Session

	if err := json.NewDecoder(file).Decode(&Response); err != nil {
		return nil, err
	}

	return Response, nil
}

func DumpJson(data []byte, path string) error {
	err := os.WriteFile(path, data, 0644)

	if err != nil {
		return err
	}

	return nil
}
