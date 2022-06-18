package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Story struct {
	arcs []StoryArc
}

func (s *Story) Load(filepath string) error {

	jsonData, err := s.getJSONData(filepath)
	if err != nil {
		return err
	}

	for key, value := range jsonData {
		arc := StoryArc{}
		arc.Load(key, value.(map[string]interface{}))
		s.arcs = append(s.arcs, arc)
	}

	return nil
}

func (s *Story) getJSONData(filepath string) (map[string]interface{}, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("Cannot find json file, error: ", err.Error())
		return nil, err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Cannot read json file, error: ", err.Error())
		return nil, err
	}

	var data interface{}
	er := json.Unmarshal(bytes, &data)
	if er != nil {
		log.Println("Cannot unmarshal json file, error: ", er.Error())
		return nil, er
	}

	return data.(map[string]interface{}), nil
}

func (s *Story) GetArc(key string) (*StoryArc, error) {
	for _, arc := range s.arcs {
		if arc.Identifier == key {
			return &arc, nil
		}
	}
	return nil, errors.New("Couldn't find " + key + " arc")
}
