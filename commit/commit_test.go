package commit_test

import (
	"regexp"
	"testing"

	"github.com/rafasf/story-summary/commit"
	"github.com/stretchr/testify/assert"
)

func GeneralTag() commit.Tag {
	return commit.Tag{regexp.MustCompile(`^General$`), "General"}
}

func SomeTags() []commit.Tag {
	return []commit.Tag{
		commit.Tag{regexp.MustCompile(`(US[0-9]+)\s*`), "Story"},
		commit.Tag{regexp.MustCompile(`(EFFIG-[0-9]+)\s*`), "Story"},
		commit.Tag{regexp.MustCompile(`chore:\s*`), "Chore"},
		commit.Tag{regexp.MustCompile(`chore:\s*`), "Chore"},
	}
}

func TestCleanMessageFromReturnsMessageWithoutTag(t *testing.T) {
	rawSubject := "Tag1 Add other to D"

	re := regexp.MustCompile(`(Tag1)\s*`)
	cleanSubject := commit.CleanMessageGiven(re, rawSubject)

	assert.Equal(t, "Add other to D", cleanSubject)
}

func TestCommitFromReturnsCleanCommit(t *testing.T) {
	logEntry := "Tag1 Add other to D|The Author|4ddf81f"

	tagMatch := commit.TagMatch{"Tag1", commit.Tag{regexp.MustCompile(`(Tag1)\s*`), "Story"}}
	actualCommit := commit.CommitFrom(tagMatch, logEntry, "|", commit.CleanMessage(tagMatch.Pattern()))
	expectedCommit := commit.Commit{
		"Add other to D",
		"The Author",
		"4ddf81f",
		tagMatch,
	}

	assert.Equal(t, expectedCommit, actualCommit)
}

func TestCommitsFromReturnsCommits(t *testing.T) {
	log := []string{
		"Add other to d|Bob|4ddf81f",
		"US234 Add new to d|John|a319639",
		"EFFIG-401 First from other tracker|Mary|f861b45",
	}

	commits := commit.CommitsFrom(log, SomeTags(), "|")
	expectedCommits := []commit.Commit{
		commit.Commit{"Add other to d", "Bob", "4ddf81f", commit.TagMatch{"General", GeneralTag()}},
		commit.Commit{"Add new to d", "John", "a319639", commit.TagMatch{"US234", SomeTags()[0]}},
		commit.Commit{"First from other tracker", "Mary", "f861b45", commit.TagMatch{"EFFIG-401", SomeTags()[1]}},
	}

	assert.Equal(t, expectedCommits[0].Tag.Description(), commits[0].Tag.Description())
	assert.Equal(t, expectedCommits[1].Tag.Description(), commits[1].Tag.Description())
	assert.Equal(t, expectedCommits[2].Tag.Description(), commits[2].Tag.Description())
}

func TestTagOrDefaultGivenReturnsTagMatch(t *testing.T) {
	tags := SomeTags()

	match := commit.TagOrDefaultGiven(tags, "US123 blah|other|things")
	expectedTag := commit.TagMatch{"US123", tags[0]}

	assert.Equal(t, expectedTag, match)
}

func TestTagOrDefaultGivenReturnsGeneralTagMatchWhenNoTag(t *testing.T) {
	tags := SomeTags()

	match := commit.TagOrDefaultGiven(tags, "blah|other|things")
	expectedTag := commit.TagMatch{"General", GeneralTag()}

	assert.Equal(t, expectedTag, match)
}

func TestByTagsReturnsCommitsGroupedByTagValue(t *testing.T) {
	commits := []commit.Commit{
		commit.Commit{"Add other to d", "Bob", "4ddf81f", commit.TagMatch{"General", GeneralTag()}},
		commit.Commit{"Add new to d", "John", "a319639", commit.TagMatch{"US234", SomeTags()[0]}},
		commit.Commit{"First from other tracker", "Mary", "f861b45", commit.TagMatch{"EFFIG-401", SomeTags()[1]}},
	}

	groups := commit.ByTag(commits)

	assert.Equal(t, groups["US234"][0].Subject, "Add new to d")
	assert.Equal(t, groups["EFFIG-401"][0].Subject, "First from other tracker")
	assert.Equal(t, groups["General"][0].Subject, "Add other to d")
}
