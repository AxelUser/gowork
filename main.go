package main

import (
	"flag"
	"strings"

	"github.com/AxelUser/gowork/parser"
)

func main() {
	usecase := flag.String("case", "web", "Define use-case for application: 'parser', 'train', 'web'")
	config := flag.String("config", "", "Define path to parser`s config file")
	flag.Parse()

	switch use := strings.ToLower(*usecase); use {
	case "parser":
		parser.Start(*config)
	}
}
