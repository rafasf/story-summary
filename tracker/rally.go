package tracker

import (
	"regexp"
)

// Rally holds the definition of Rally's tracker.
type Rally struct {
	Info Tracker
}

// StoryFor returns a Story by looking up the info in the tracker.
func (t Rally) StoryFor(storyID string) Story {
	return Story{storyID, "The Cool Summary"}
}

// AllPatterns returns the combined regex of all tracker patterns.
func (t Rally) AllPatterns() *regexp.Regexp {
	return CombinedPatternsGiven(t.Info.Patterns)
}

// Details returns the tracker information.
func (t Rally) Details() Tracker {
	return t.Info
}
