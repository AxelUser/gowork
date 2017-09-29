package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/AxelUser/gowork/errors"
	"github.com/AxelUser/gowork/models"
	"github.com/AxelUser/gowork/parser/loader"
)

func save() {

}

func getParserConfig(configPath string) (*models.ParserConfig, error) {
	raw, fsErr := ioutil.ReadFile(configPath)
	if fsErr != nil {
		return nil, errors.NewConfigLoadError(fmt.Sprintf("Could not read file '%s'", configPath), fsErr)
	}

	var config models.ParserConfig
	jsonErr := json.Unmarshal(raw, &config)
	if jsonErr != nil {
		return nil, errors.NewConfigLoadError(fmt.Sprintf("Could not unmarshal json '%s'", configPath), jsonErr)
	}

	return &config, nil
}

//Start loading data
func Start(configPath string) {
	log.Print("Parser started.")
	config, confErr := getParserConfig(configPath)
	if confErr != nil {
		log.Fatal(confErr)
	} else {
		log.Print("Config loaded")
	}
	loader.Load(*config)
}
