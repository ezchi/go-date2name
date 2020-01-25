package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var inFileName string

	flag.StringVar(&inFileName, "f", "", "file name")

	flag.Parse()

	if inFileName == "" {
		log.Fatal("File name is missing")
	}

	inFileName, _ = removeSpaces(inFileName)

	_, err := getTimeFromName(inFileName)

	if err == nil {
		os.Exit(0)
	}

	s, err := getModifiedTime(inFileName)

	if err != nil {
		log.Fatalf("can not get modified name of %s: %v", inFileName, err)
	}

	b := filepath.Base(inFileName)
	d := filepath.Dir(inFileName)

	outFileName := filepath.Join(d, fmt.Sprintf("%s-%s", s, b))

	err = rename(inFileName, outFileName)

	if err != nil {
		log.Fatal(err)
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func rename(oldName, newName string) error {
	if isExist(newName) {
		return fmt.Errorf("file %s already exist", newName)
	}

	return os.Rename(oldName, newName)
}

func removeSpaces(path string) (string, error) {
	s := strings.Split(path, " ")

	f := strings.Join(s, "_")

	return f, rename(path, f)
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
