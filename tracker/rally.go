package tracker

import (
	"regexp"
)

type Rally struct {
	Info Tracker
}

func (t Rally) StoryFor(storyId string) Story {
	return Story{storyId, "The Cool Summary"}
}

func (t Rally) AllPatterns() *regexp.Regexp {
	return CombinedPatternsGiven(t.Info.Patterns)
}

func (t Rally) Details() Tracker {
	return t.Info
}
