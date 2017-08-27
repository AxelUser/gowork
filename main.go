package main

import (
	"flag"
	"strings"

	"gowork/parser"
)

func main() {
	usecase := flag.String("case", "web", "Define use-case for application: 'load', 'train', 'web'")
	config := flag.String("config", "", "Define path to parser`s config file")
	flag.Parse()

	switch use := strings.ToLower(*usecase); use {
	case "parser":
		startParsing(*config)
	}
}

func startParsing(configPath string) {
	parser.Start(configPath)
}
