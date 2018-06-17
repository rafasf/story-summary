package tracker

import (
	"regexp"
)

// Jira holds the definition of Jira's tracker.
type Jira struct {
	Info Tracker
}

// StoryFor returns a Story by looking up the info in the tracker.
func (t Jira) StoryFor(storyID string) Story {
	return Story{storyID, "The Cool Summary"}
}

// AllPatterns returns the combined regex of all tracker patterns.
func (t Jira) AllPatterns() *regexp.Regexp {
	return CombinedPatternsGiven(t.Info.Patterns)
}

// Details returns the tracker information.
func (t Jira) Details() Tracker {
	return t.Info
}
