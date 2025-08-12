package token_test

import (
	"fmt"
	"testing"

	"github.com/perdue/syt/pkg/token"
)

func TestTokenGet(t *testing.T) {
	token, err := token.Get()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	fmt.Printf("%+v", token)
}
