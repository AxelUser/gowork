package main

import (
	"flag"
	"strings"

	"github.com/AxelUser/gowork/parser"
)

func main() {
	usecase := flag.String("case", "web", "Define use-case for application: 'parser', 'train', 'web'")
	config := flag.String("config", "", "Define path to parser`s config file")
	ontology := flag.String("ontology", "", "Define path to app's ontology of skills")
	flag.Parse()

	switch use := strings.ToLower(*usecase); use {
	case "parser":
		parser.Start(*config, *ontology)
	}
}
