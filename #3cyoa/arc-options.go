package main

type ArcOptions struct {
	Number int
	Text   string
	Arc    string
}

func (ao *ArcOptions) Load(number int, data map[string]interface{}) {
	ao.Number = number
	ao.Text = data["text"].(string)
	ao.Arc = data["arc"].(string)
}
