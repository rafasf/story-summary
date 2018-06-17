package tracker

import (
	"regexp"
	"strings"
)

// Story represents a unit of work.
type Story struct {
	Identifier string
	Summary    string
}

// LookupTracker exposes what is necessary for each tracker to implement.
type LookupTracker interface {
	AllPatterns() *regexp.Regexp
	StoryFor(storyID string) Story
	Details() Tracker
}

// Tracker represents the basic information of an issue tracker.
type Tracker struct {
	Name     string
	BaseURL  string
	Patterns []*regexp.Regexp
}

// CombinedPatternsGiven returns one regex considering all given patterns.
func CombinedPatternsGiven(patterns []*regexp.Regexp) *regexp.Regexp {
	var allPatterns []string
	for _, pattern := range patterns {
		allPatterns = append(allPatterns, pattern.String())
	}
	return regexp.MustCompile(strings.Join(allPatterns, "|"))
}

// Given finds the right tracker for lookup given a story issue.
func Given(tag string, trackers []LookupTracker) LookupTracker {
	for _, tracker := range trackers {
		match := tracker.AllPatterns().FindAllStringSubmatch(tag, -1)
		if len(match) > 0 {
			return tracker
		}
	}
	return nil
}
