package main

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestDefaultTagOrReturnsGeneralWhenNoTag(t *testing.T) {
	var emptyMatch [][]string

	defaultTag := DefaultTagOr(emptyMatch)

	assert.Equal(t, "General", defaultTag)
}

func TestDefaultTagOrReturnsFirstMatchWithoutSpaceWhenPresent(t *testing.T) {
	matched := [][]string{
		[]string{"TheTag "},
	}

	aTag := DefaultTagOr(matched)

	assert.Equal(t, "TheTag", aTag)
}

func TestCleanMessageFromReturnsMessageWithoutTag(t *testing.T) {
	rawSubject := "Tag1 Add other to D"

	re := regexp.MustCompile(`(Tag1)\s*`)
	cleanSubject := CleanMessageGiven(re, rawSubject)

	assert.Equal(t, "Add other to D", cleanSubject)
}

func TestCommitFromReturnsCleanCommit(t *testing.T) {
	logEntry := "Tag1 Add other to D|The Author|4ddf81f"

	re := regexp.MustCompile(`(Tag1)\s*`)
	commit := CommitFrom("Tag1", logEntry, "|", CleanMessage(re))
	expectedCommit := Commit{
		"Add other to D",
		"The Author",
		"4ddf81f",
		"Tag1",
	}

	assert.Equal(t, expectedCommit, commit)
}

func TestCommitsFromReturnsCommits(t *testing.T) {
	log := []string{
		"Add other to d|Bob|4ddf81f",
		"US234 Add new to d|John|a319639",
		"EFFIG-401 First from other tracker|Mary|f861b45",
	}

	re := regexp.MustCompile(`(US[0-9]+)\s*|(EFFIG-[0-9]+)\s*`)
	commits := CommitsFrom(log, re, "|")
	expectedCommits := []Commit{
		Commit{"Add other to d", "Bob", "4ddf81f", "General"},
		Commit{"Add new to d", "John", "a319639", "US234"},
		Commit{"First from other tracker", "Mary", "f861b45", "EFFIG-401"},
	}

	assert.Equal(t, expectedCommits, commits)
}
