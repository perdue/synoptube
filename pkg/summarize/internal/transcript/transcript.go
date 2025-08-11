package transcript

import (
	"fmt"
	"log"
	"strings"

	"github.com/chand1012/yt_transcript"
)

func FetchTranscript(videoURL string) (string, error) {
	videoId, err := yt_transcript.GetVideoID(videoURL)
	if err != nil {
		return "", fmt.Errorf("error getting Id for %s: %v", videoURL, err)
	}

	responses, title, err := yt_transcript.FetchTranscript(videoId, "en", "US")
	if err != nil {
		return "", fmt.Errorf("error fetching YouTube transcript for %s: %v", videoId, err)
	}

	var fullTranscript strings.Builder
	for _, t := range responses {
		fullTranscript.WriteString(t.Text)
		fullTranscript.WriteString(" ")
	}
	transcriptText := fullTranscript.String()

	if len(transcriptText) == 0 {
		return "", fmt.Errorf("no transcript found for video %s", videoId)
	}

	log.Printf("Fetched transcript for %s (%s) (first 200 chars): %s...\n", title, videoId, transcriptText[:min(len(transcriptText), 200)])

	return transcriptText, nil
}
