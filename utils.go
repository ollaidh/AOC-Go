package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func getInputsFolderPath() (string, error) {
	inputsFolderRootPath := os.Getenv("AOC_INPUTS_FOLDER")
	if inputsFolderRootPath == "" {
		return "", errors.New("No AOC_INPUTS_FOLDER environment variable is found!")
	}
	return inputsFolderRootPath, nil
}

func getInputFilePaths(date Date, inputRootPath string) []string {
	fileName := fmt.Sprintf("input_day_%02d.txt", date.day)
	fileNames := []string{fileName} // TODO no glob so far - ADD!
	for i := range fileNames {
		fileNames[i] = filepath.Join(inputRootPath, strconv.Itoa(date.year), fileNames[i])
	}
	return fileNames
}

func getDataFromFile(inputPath string) []string {
	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}
