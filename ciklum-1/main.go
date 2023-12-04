package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	print("hello world")
}

type domain struct {
	name       string
	count      int
	subDomains map[string]domain
}

func calculateClicksByDomain(inputFileName string, outputFileName string) {

	//  Implement the logic for calculate clicks
	var tld domain
	tld.subDomains = make(map[string]domain)

	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		saveStats(scanner.Text(), tld)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	//save results to in domain order
	//TODO
}

func saveStats(name string,d domain){
	names := strings.Split(name, ".")
	d.name = name
	d.count++
	saveStats(names[1:],d) //conver []string to string
}
