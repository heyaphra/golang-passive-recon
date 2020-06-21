package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexsasharegan/dotenv"
	"github.com/spidercatnat/golang-passive-recon/shodan"
)

func getEnv() {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	getEnv()
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d %s\n", host.IPString, host.Port, host.Hostnames)
	}
}
