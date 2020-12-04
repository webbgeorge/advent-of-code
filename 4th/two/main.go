package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	validCount := 0

	scanner := bufio.NewScanner(f)
	scanner.Split(splitByDoubleNewLine)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)

		passport := make(map[string]string)
		for _, field := range strings.Fields(line) {
			passport[field[0:3]] = field[4:]
		}
		if validatePassport(passport) {
			validCount++
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(validCount)
}

var validateBYR = regexp.MustCompile(`^(19[2-9]\d|200[0-2])$`) //  four digits; at least 1920 and at most 2002.
var validateIYR = regexp.MustCompile(`^(201\d|2020)$`) // four digits; at least 2010 and at most 2020.
var validateEYR = regexp.MustCompile(`^(202\d|2030)$`) // four digits; at least 2020 and at most 2030.
var validateHGT = regexp.MustCompile(`^((59|6\d|7[0-6])in|1([5-8]\d|9[0-3])cm)$`)
var validateHCL = regexp.MustCompile(`^#[0-9a-f]{6}$`) // a # followed by exactly six characters 0-9 or a-f.
var validateECL = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`) // one of: amb blu brn gry grn hzl oth.
var validatePID = regexp.MustCompile(`^\d{9}$`)

func validatePassport(passport map[string]string) bool {
	if !validateBYR.MatchString(passport["byr"]) {
		return false
	}
	if !validateIYR.MatchString(passport["iyr"]) {
		return false
	}
	if !validateEYR.MatchString(passport["eyr"]) {
		return false
	}
	if !validateHGT.MatchString(passport["hgt"]) {
		return false
	}
	if !validateHCL.MatchString(passport["hcl"]) {
		return false
	}
	if !validateECL.MatchString(passport["ecl"]) {
		return false
	}
	if !validatePID.MatchString(passport["pid"]) {
		return false
	}
	return true
}

func splitByDoubleNewLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\n','\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i+1]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
