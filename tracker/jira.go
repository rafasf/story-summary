package tracker

import (
	"regexp"
)

type Jira struct {
	Info Tracker
}

func (t Jira) StoryFor(storyId string) Story {
	return Story{storyId, "The Cool Summary"}
}

func (t Jira) AllPatterns() *regexp.Regexp {
	return CombinedPatternsGiven(t.Info.Patterns)
}

func (t Jira) Details() Tracker {
	return t.Info
}
