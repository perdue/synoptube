package summarize

import "github.com/perdue/synoptube/pkg/summarize/internal/transcript"

func ProcessVideo() error {
	if _, err := transcript.FetchTranscript("https://youtu.be/O83iXjT2r8A"); err != nil {
		return err
	}

	return nil
}
