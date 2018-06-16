package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Commit struct {
	subject string
	author  string
	hash    string
	tag     string
}

type CleanMsg func(string) string

func DefaultTagOr(match [][]string) string {
	if len(match) > 0 {
		return strings.TrimSpace(match[0][0])
	} else {
		return "General"
	}
}

func CleanMessage(re *regexp.Regexp) func(string) string {
	return func(subject string) string {
		return CleanMessageGiven(re, subject)
	}
}

func CleanMessageGiven(re *regexp.Regexp, subject string) string {
	return re.ReplaceAllString(subject, "")
}

func CommitFrom(tag string, logEntry string, separator string, cleaner CleanMsg) Commit {
	parts := strings.Split(logEntry, separator)
	subject, author, hash := parts[0], parts[1], parts[2]
	return Commit{cleaner(subject), author, hash, tag}
}

func CommitsFrom(logEntries []string, tag_re *regexp.Regexp, separator string) []Commit {
	var commits []Commit
	for _, logEntry := range logEntries {
		possibleTag := tag_re.FindAllStringSubmatch(logEntry, -1)
		commits = append(
			commits,
			CommitFrom(
				DefaultTagOr(possibleTag),
				logEntry,
				separator,
				CleanMessage(tag_re)))
	}
	return commits
}

func main() {
	logEntries := []string{
		"Add other to d|Rafael Ferreira|4ddf81f",
		"Add new to d|Rafael Ferreira|a319639",
		"EFFIG-401 First from other tracker|Rafael Ferreira|f861b45",
		"US11713 Aenean vestibulum ipsum et|Rafael Ferreira|cbe86a1",
		"US13791 Suspendisse dignissim hendrerit porttitor|Rafael Ferreira|cd4730f",
		"chore: Initial commit|Rafael Ferreira|a041e85"}

	re, _ := regexp.Compile(`(US[0-9]+)\s*|(EFFIG-[0-9]+)\s*`)
	commits := CommitsFrom(logEntries, re, "|")

	for _, commit := range commits {
		fmt.Println(commit)
	}
}
