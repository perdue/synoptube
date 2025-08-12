package credentials

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func ConfigFromSecrets(fileName string) (*oauth2.Config, error) {
	var bytes []byte
	var err error
	if bytes, err = ReadFile(fileName); err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(bytes, youtube.YoutubeReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("problem parsing secrets file to config: %s", err)
	}

	return config, nil
}

func ReadFile(fileName string) ([]byte, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("problem reading secrets file '%s': %s", fileName, err)
	}

	return bytes, nil
}

// return bytes, nil
