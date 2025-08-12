package token

import (
	"fmt"

	credentials "github.com/perdue/syt/pkg/token/internal/credentials"
	"golang.org/x/oauth2"
)

func Get() (*oauth2.Token, error) {
	cacheFile, err := credentials.CacheFile(".syt", "credentials.json")
	if err != nil {
		err = fmt.Errorf("problem creating credentials cache file: %s", err)

		return nil, err
	}

	fmt.Println(cacheFile)

	creds, err := credentials.FromFile(cacheFile)
	if err != nil {
		//   tok = getTokenFromWeb(config)
		//   saveToken(cacheFile, tok)
	}

	fmt.Println(creds)

	return nil, nil
}
