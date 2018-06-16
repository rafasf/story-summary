package main

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func GeneralTag() Tag {
	return Tag{regexp.MustCompile(`^General$`), "General"}
}

func SomeTags() []Tag {
	return []Tag{
		Tag{regexp.MustCompile(`(US[0-9]+)\s*`), "Story"},
		Tag{regexp.MustCompile(`(EFFIG-[0-9]+)\s*`), "Story"},
		Tag{regexp.MustCompile(`chore:\s*`), "Chore"},
		Tag{regexp.MustCompile(`chore:\s*`), "Chore"},
	}
}

func TestCleanMessageFromReturnsMessageWithoutTag(t *testing.T) {
	rawSubject := "Tag1 Add other to D"

	re := regexp.MustCompile(`(Tag1)\s*`)
	cleanSubject := CleanMessageGiven(re, rawSubject)

	assert.Equal(t, "Add other to D", cleanSubject)
}

func TestCommitFromReturnsCleanCommit(t *testing.T) {
	logEntry := "Tag1 Add other to D|The Author|4ddf81f"

	tagMatch := TagMatch{"Tag1", Tag{regexp.MustCompile(`(Tag1)\s*`), "Story"}}
	commit := CommitFrom(tagMatch, logEntry, "|", CleanMessage(tagMatch.Pattern()))
	expectedCommit := Commit{
		"Add other to D",
		"The Author",
		"4ddf81f",
		tagMatch,
	}

	assert.Equal(t, expectedCommit, commit)
}

func TestCommitsFromReturnsCommits(t *testing.T) {
	log := []string{
		"Add other to d|Bob|4ddf81f",
		"US234 Add new to d|John|a319639",
		"EFFIG-401 First from other tracker|Mary|f861b45",
	}

	commits := CommitsFrom(log, SomeTags(), "|")
	expectedCommits := []Commit{
		Commit{"Add other to d", "Bob", "4ddf81f", TagMatch{"General", GeneralTag()}},
		Commit{"Add new to d", "John", "a319639", TagMatch{"US234", SomeTags()[0]}},
		Commit{"First from other tracker", "Mary", "f861b45", TagMatch{"EFFIG-401", SomeTags()[1]}},
	}

	assert.Equal(t, expectedCommits[0].tag.Description(), commits[0].tag.Description())
	assert.Equal(t, expectedCommits[1].tag.Description(), commits[1].tag.Description())
	assert.Equal(t, expectedCommits[2].tag.Description(), commits[2].tag.Description())
}

func TestTagOrDefaultGivenReturnsTagMatch(t *testing.T) {
	tags := SomeTags()

	match := TagOrDefaultGiven(tags, "US123 blah|other|things")
	expectedTag := TagMatch{"US123", tags[0]}

	assert.Equal(t, expectedTag, match)
}

func TestTagOrDefaultGivenReturnsGeneralTagMatchWhenNoTag(t *testing.T) {
	tags := SomeTags()

	match := TagOrDefaultGiven(tags, "blah|other|things")
	expectedTag := TagMatch{"General", GeneralTag()}

	assert.Equal(t, expectedTag, match)
}
