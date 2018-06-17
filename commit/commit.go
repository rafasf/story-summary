package commit

import (
	"regexp"
	"strings"
)

type Commit struct {
	Subject string
	Author  string
	Hash    string
	Tag     TagMatch
}

type Tag struct {
	Pattern     *regexp.Regexp
	Description string
}

type TagMatch struct {
	Value  string
	Source Tag
}

func (tM TagMatch) Description() string {
	return tM.Source.Description
}

func (tM TagMatch) Pattern() *regexp.Regexp {
	return tM.Source.Pattern
}

type CleanMsg func(string) string

func TagOrDefaultGiven(tags []Tag, entry string) TagMatch {
	for _, tag := range tags {
		possibleMatch := tag.Pattern.FindAllStringSubmatch(entry, -1)
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
		value := commit.Tag.Value
		commitsByTag[value] = append(commitsByTag[value], commit)
	}
	return commitsByTag
}
