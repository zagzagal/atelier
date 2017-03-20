package main

import (
	"bufio"
	"log"
	"os"
	//"fmt"
	ad "bitbucket.org/zagzagal/AtelierComplete/AtelierData"
	"regexp"
)

func loadData(fName string) *ad.AtelierData {
	file, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var validLine = regexp.MustCompile(`"(.*)" -> "(.*)"`)
	data := ad.NewAtelier()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := validLine.FindStringSubmatch(scanner.Text())
		if line != nil {
			data.AddPath(line[1], line[2])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func writeData(fName string, data string) {
	f, err := os.Create(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(data)
	f.Sync()
}
