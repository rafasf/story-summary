package tracker

import (
	"regexp"
	"strings"
)

type Story struct {
	Identifier string
	Summary    string
}

type LookupTracker interface {
	AllPatterns() *regexp.Regexp
	StoryFor(storyID string) Story
	Details() Tracker
}

type Tracker struct {
	Name     string
	BaseUrl  string
	Patterns []*regexp.Regexp
}

func CombinedPatternsGiven(patterns []*regexp.Regexp) *regexp.Regexp {
	var allPatterns []string
	for _, pattern := range patterns {
		allPatterns = append(allPatterns, pattern.String())
	}
	return regexp.MustCompile(strings.Join(allPatterns, "|"))
}

func TrackerGiven(tag string, trackers []LookupTracker) LookupTracker {
	for _, tracker := range trackers {
		match := tracker.AllPatterns().FindAllStringSubmatch(tag, -1)
		if len(match) > 0 {
			return tracker
		}
	}
	return nil
}
