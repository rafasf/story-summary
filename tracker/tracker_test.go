package tracker_test

import (
	"regexp"
	"testing"

	"github.com/rafasf/story-summary/tracker"
	"github.com/stretchr/testify/assert"
)

func JiraTracker() tracker.LookupTracker {
	return tracker.Jira{
		Info: tracker.Tracker{
			Name:    "Jira",
			BaseURL: "http://jira.com",
			Patterns: []*regexp.Regexp{
				regexp.MustCompile(`(EFFIG-[0-9])\s*`),
			},
		},
	}
}

func RallyTracker() tracker.LookupTracker {
	return tracker.Rally{
		Info: tracker.Tracker{
			Name:    "Rally",
			BaseURL: "http://rally.com",
			Patterns: []*regexp.Regexp{
				regexp.MustCompile(`(US[0-9])\s*`),
				regexp.MustCompile(`(DE[0-9])\s*`),
			},
		},
	}
}

func TestGivenReturnsProperTrackerForStorySummaryLookup(t *testing.T) {
	trackers := []tracker.LookupTracker{JiraTracker(), RallyTracker()}

	assert.Equal(t, tracker.Given("EFFIG-401", trackers).Details().Name, "Jira")
	assert.Equal(t, tracker.Given("US123", trackers).Details().Name, "Rally")
	assert.Equal(t, tracker.Given("DE123", trackers).Details().Name, "Rally")
}
