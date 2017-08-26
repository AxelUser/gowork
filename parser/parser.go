package parser

import (
	"encoding/json"
	"fmt"
	"gowork/errors"
	"gowork/models"
	"io/ioutil"
)

func load() {

}

func save() {

}

func getParserConfig(configPath string) (*models.ParserConfig, error) {
	raw, fsErr := ioutil.ReadFile(configPath)
	if fsErr != nil {
		return nil, errors.NewConfigLoadError(fmt.Sprintf("Could not read file '%s'", configPath), fsErr)
	}

	var config *models.ParserConfig
	jsonErr := json.Unmarshal(raw, config)
	if jsonErr != nil || config == nil {
		return nil, errors.NewConfigLoadError(fmt.Sprintf("Could not unmarshal json '%s'", configPath), jsonErr)
	}

	return config, nil
}

//Start loading data
func Start(configPath string) {

}
