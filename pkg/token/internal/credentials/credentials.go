package credentials

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/oauth2"
)

// CacheFile generates credential file path/filename
// in the user's home directory.
func CacheFile(storageDir string, fileName string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, storageDir)
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape(fileName)), err
}

func FromFile(fileName string) (*oauth2.Token, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("problem opening '%s' file: %s", fileName, err)
	}
	defer f.Close()

	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	if err != nil {
		return nil, fmt.Errorf("problem decoding credentials in '%s': %s", fileName, err)
	}

	return t, nil
}

func FromWeb(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, fmt.Errorf("problem reading authorization code: %s", err)
	}

	t, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("problem getting authorization code: %s", err)
	}

	return t, nil
}
