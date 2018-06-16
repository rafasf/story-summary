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
	tag     TagMatch
}

type Tag struct {
	pattern     *regexp.Regexp
	description string
}

type TagMatch struct {
	value  string
	source Tag
}

func (tM TagMatch) Description() string {
	return tM.source.description
}

func (tM TagMatch) Pattern() *regexp.Regexp {
	return tM.source.pattern
}

type CleanMsg func(string) string

func TagOrDefaultGiven(tags []Tag, entry string) TagMatch {
	for _, tag := range tags {
		possibleMatch := tag.pattern.FindAllStringSubmatch(entry, -1)
		if len(possibleMatch) > 0 {
			return TagMatch{strings.TrimSpace(possibleMatch[0][0]), tag}
		}
	}
	return TagMatch{"General", Tag{regexp.MustCompile(`^General$`), "General"}}
}

func CleanMessage(re *regexp.Regexp) func(string) string {
	return func(subject string) string {
		return CleanMessageGiven(re, subject)
	}
}

func CleanMessageGiven(re *regexp.Regexp, subject string) string {
	return re.ReplaceAllString(subject, "")
}

func CommitFrom(tagMatch TagMatch, logEntry string, separator string, cleaner CleanMsg) Commit {
	parts := strings.Split(logEntry, separator)
	subject, author, hash := parts[0], parts[1], parts[2]
	return Commit{cleaner(subject), author, hash, tagMatch}
}

func CommitsFrom(logEntries []string, tags []Tag, separator string) []Commit {
	var commits []Commit
	for _, logEntry := range logEntries {
		tagMatch := TagOrDefaultGiven(tags, logEntry)
		commits = append(
			commits,
			CommitFrom(
				tagMatch,
				logEntry,
				separator,
				CleanMessage(tagMatch.Pattern())))
	}
	return commits
}

func ByTag(commits []Commit) map[string][]Commit {
	commitsByTag := make(map[string][]Commit)
	for _, commit := range commits {
		value := commit.tag.value
		commitsByTag[value] = append(commitsByTag[value], commit)
	}
	return commitsByTag
}

type Tracker struct {
	name     string
	baseUrl  string
	patterns *[]regexp.Regexp
}

func main() {
	logEntries := []string{
		"Add other to d|Rafael Ferreira|4ddf81f",
		"Add new to d|Rafael Ferreira|a319639",
		"EFFIG-401 First from other tracker|Rafael Ferreira|f861b45",
		"US11713 Aenean vestibulum ipsum et|Rafael Ferreira|cbe86a1",
		"US13791 Suspendisse dignissim hendrerit porttitor|Rafael Ferreira|cd4730f",
		"chore: Initial commit|Rafael Ferreira|a041e85"}

	tags := []Tag{
		Tag{regexp.MustCompile(`(US[0-9]+)\s*`), "Story"},
		Tag{regexp.MustCompile(`(EFFIG-[0-9]+)\s*`), "Story"},
		Tag{regexp.MustCompile(`chore:\s*`), "Chore"},
	}

	trackers := []Tracker{
		Tracker{
			"Jira",
			"http://jira.com",
			[]regexp.Regexp{
				regexp.MustCompile(`(EFFIG-[0-9])\s*`),
			},
		},
		Tracker{
			"Rally",
			"http://rally.com",
			[]regexp.Regexp{
				regexp.MustCompile(`(US[0-9])\s*`),
				regexp.MustCompile(`(DE[0-9])\s*`),
			},
		},
	}

	commits := CommitsFrom(logEntries, tags, "|")

	for _, commit := range commits {
		fmt.Println(commit)
	}

	for k, v := range ByTag(commits) {
		fmt.Printf("key: %s\n", k)
		for _, c := range v {
			fmt.Println(c)
		}

		fmt.Println("")
	}
}
