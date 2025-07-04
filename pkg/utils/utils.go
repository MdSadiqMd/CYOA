package utils

import (
	"encoding/json"
	"io"

	"github.com/MdSadiqMd/CYOA/pkg/schema"
)

func JsonStory(r io.Reader) (story schema.Story, error error) {
	d := json.NewDecoder(r)
	var Story schema.Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return Story, nil
}
