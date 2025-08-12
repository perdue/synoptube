package token

import (
	"context"
	"fmt"

	credentials "github.com/perdue/syt/pkg/token/internal/credentials"
	"golang.org/x/oauth2"
)

func Get(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	cacheFile, err := credentials.CacheFile(".syt", "credentials.json")
	if err != nil {
		err = fmt.Errorf("problem creating credentials cache file: %s", err)

		return nil, err
	}

	fmt.Println(cacheFile)

	creds, err := credentials.FromFile(cacheFile)
	if err == nil {
		return creds, nil
	}

	creds, err = credentials.FromWeb(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("problem getting credentials: %s", err)
	}
	//   saveToken(cacheFile, tok)

	fmt.Println(creds)

	return creds, nil
}
