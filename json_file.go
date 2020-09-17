package nhl_stats_api

import (
	"fmt"
	"io/ioutil"
	"os"
)

func getJsonFile(relativePath string) ([]byte, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, relativePath))
	if err != nil {
		return nil, err
	}

	return file, nil
}
