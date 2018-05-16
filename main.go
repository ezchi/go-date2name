package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	inFileName string
)

func main() {
	flag.StringVar(&inFileName, "f", "", "file name")

	flag.Parse()

	_, err := getTimeFromName(inFileName)

	if err == nil {
		os.Exit(0)
	}

	s, err := getModifiedTime(inFileName)

	if err != nil {
		log.Fatalf("can not get modified name of %s: %v", inFileName, err)
	}

	outFileName := fmt.Sprintf("%s_%s", s, inFileName)

	_, err = os.Stat(outFileName)

	if !os.IsNotExist(err) {
		log.Fatalf("File %s is already exist", outFileName)
	}

	err = os.Rename(inFileName, outFileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func getTimeFromName(path string) (string, error) {
	s := strings.SplitN(path, "_", 2)

	_, err := time.Parse(time.RFC3339, s[0])

	if err != nil {
		return "", err
	}

	return s[0], nil
}

func getModifiedTime(path string) (string, error) {
	info, err := os.Lstat(path)

	if err != nil {
		return "", err
	}

	return info.ModTime().Format(time.RFC3339), nil
}
