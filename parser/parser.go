package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/AxelUser/gowork/parser/normalizer"

	"github.com/AxelUser/gowork/errors"
	"github.com/AxelUser/gowork/models/configs"
	"github.com/AxelUser/gowork/parser/loader"
)

func save() {

}

func getParserConfig(configPath string) (*configs.ParserConfig, error) {
	raw, fsErr := ioutil.ReadFile(configPath)
	if fsErr != nil {
		return nil, errors.NewConfigLoadError(fmt.Sprintf("Could not read file '%s'", configPath), fsErr)
	}

	var config configs.ParserConfig
	jsonErr := json.Unmarshal(raw, &config)
	if jsonErr != nil {
		return nil, errors.NewConfigLoadError(fmt.Sprintf("Could not unmarshal json '%s'", configPath), jsonErr)
	}

	return &config, nil
}

func getOntology(ontologyPath string) ([]configs.OntologyData, error) {
	raw, fsErr := ioutil.ReadFile(ontologyPath)
	if fsErr != nil {
		return nil, errors.NewConfigLoadError(fmt.Sprintf("Could not read file '%s'", ontologyPath), fsErr)
	}

	var ontologyInfos []configs.OntologyData
	jsonErr := json.Unmarshal(raw, &ontologyInfos)
	if jsonErr != nil {
		return nil, errors.NewConfigLoadError(fmt.Sprintf("Could not unmarshal json '%s'", ontologyPath), jsonErr)
	}

	return ontologyInfos, nil
}

//Start loading and normalizing data
func Start(configPath string, ontologyPath string) {
	log.Print("Parser started")
	config, confErr := getParserConfig(configPath)
	if confErr != nil {
		log.Fatal(confErr)
	} else {
		log.Print("Config loaded")
	}
	rawData, err := loader.Load(*config)

	if err != nil {
		log.Fatal(err)
	}

	ontologyInfos, confErr := getOntology(ontologyPath)
	if confErr != nil {
		log.Fatal(confErr)
	} else {
		log.Print("Ontology loaded")
	}

	normalizer.NormalizeRawData(ontologyInfos, rawData)

}
