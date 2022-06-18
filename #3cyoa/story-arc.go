package main

import "bytes"

type StoryArc struct {
	Identifier string
	Title      string
	Paragraph  string
	Options    []*ArcOptions
}

func (sa *StoryArc) Load(key string, data map[string]interface{}) {
	var buffer bytes.Buffer
	paragraphs := data["story"].([]interface{})

	for _, p := range paragraphs {
		buffer.WriteString(p.(string) + "\r\n")
	}

	sa.Identifier = key
	sa.Title = data["title"].(string)
	sa.Paragraph = buffer.String()
}
