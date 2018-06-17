package commit

import (
	"regexp"
	"strings"
)

// Commit represents a clean representation of a commit.
type Commit struct {
	Subject string
	Author  string
	Hash    string
	Tag     TagMatch
}

// Tag represents the different values used to group commits.
type Tag struct {
	Pattern     *regexp.Regexp
	Description string
}

// TagMatch represents a matched tag (or not) in a commit.
type TagMatch struct {
	Value  string
	Source Tag
}

// Description eturns tag's description.
func (tM TagMatch) Description() string {
	return tM.Source.Description
}

// Pattern returns tag's pattern.
func (tM TagMatch) Pattern() *regexp.Regexp {
	return tM.Source.Pattern
}

// CleanMsg is an alias for function to clean commit messages.
type CleanMsg func(string) string

// TagOrDefaultGiven returns a TagMatch when the commits contain it or
// the General tag, without value.
func TagOrDefaultGiven(tags []Tag, entry string) TagMatch {
	for _, tag := range tags {
		possibleMatch := tag.Pattern.FindAllStringSubmatch(entry, -1)
		if len(possibleMatch) > 0 {
			return TagMatch{strings.TrimSpace(possibleMatch[0][0]), tag}
		}
	}
	return TagMatch{"General", Tag{regexp.MustCompile(`^General$`), "General"}}
}

// CleanMessage holds the pattern to be used when cleaning commit message.
func CleanMessage(re *regexp.Regexp) func(string) string {
	return func(subject string) string {
		return CleanMessageGiven(re, subject)
	}
}

// CleanMessageGiven removes pattern occurences from given subject.
func CleanMessageGiven(re *regexp.Regexp, subject string) string {
	return re.ReplaceAllString(subject, "")
}

// From creates a commit from a log entry.
func From(tagMatch TagMatch, logEntry string, separator string, cleaner CleanMsg) Commit {
	parts := strings.Split(logEntry, separator)
	subject, author, hash := parts[0], parts[1], parts[2]
	return Commit{cleaner(subject), author, hash, tagMatch}
}

// CommitsFrom creates commits from the log entries provided.
func CommitsFrom(logEntries []string, tags []Tag, separator string) []Commit {
	var commits []Commit
	for _, logEntry := range logEntries {
		tagMatch := TagOrDefaultGiven(tags, logEntry)
		commits = append(
			commits,
			From(
				tagMatch,
				logEntry,
				separator,
				CleanMessage(tagMatch.Pattern())))
	}
	return commits
}

// ByTag groups commits by tag value
func ByTag(commits []Commit) map[string][]Commit {
	commitsByTag := make(map[string][]Commit)
	for _, commit := range commits {
		value := commit.Tag.Value
		commitsByTag[value] = append(commitsByTag[value], commit)
	}
	return commitsByTag
}
