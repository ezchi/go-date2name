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

	b := filepath.Base(inFileName)

	_, ok := getTimeFromName(b)

	if ok {
		os.Exit(0)
	}

	ts, err := getModifiedTime(inFileName)

	if err != nil {
		log.Fatalf("can not get modified name of %s: %v", inFileName, err)
	}

	d := filepath.Dir(inFileName)
	outFileName := filepath.Join(d, fmt.Sprintf("%s-%s", ts, b))

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

func getTimeFromName(path string) (string, bool) {
	if len(path) < 25 {
		return "", false
	}

	ts := path[:25]

	_, err := time.Parse(time.RFC3339, ts)

	if err != nil {
		return "", false
	}

	return ts, true
}

func getModifiedTime(path string) (string, error) {
	info, err := os.Lstat(path)

	if err != nil {
		return "", err
	}

	return info.ModTime().Format(time.RFC3339), nil
}
