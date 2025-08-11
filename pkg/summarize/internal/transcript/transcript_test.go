package transcript_test

import (
	"fmt"
	"testing"

	"github.com/perdue/synoptube/pkg/summarize/internal/transcript"
)

func TestFetchTranscript(t *testing.T) {
	testURL := "https://youtu.be/FZwmI0u8Qk8"
	result, err := transcript.FetchTranscript(testURL)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	fmt.Printf("%+v", result)
}
