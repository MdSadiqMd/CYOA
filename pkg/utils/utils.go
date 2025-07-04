package utils

import (
	"encoding/json"
	"io"

	"github.com/MdSadiqMd/CYOA/pkg/schema"
)

func JsonStory(r io.Reader) (schema.Story, error) {
	var story schema.Story
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}
