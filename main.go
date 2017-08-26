package main

import (
	"flag"
	"log"
	"strings"

	"gowork/parser"
)

func main() {
	usecase := flag.String("case", "web", "Define use-case for application: 'load', 'train', 'web'")
	flag.Parse()

	switch use := strings.ToLower(*usecase); use {
	case "parser":
		startParsing()
	}
}

func startParsing() {
	log.Print("Parser started.")
	parser.Start()
}
