package tracker_test

import (
	"regexp"
	"testing"

	"github.com/rafasf/story-summary/tracker"
	"github.com/stretchr/testify/assert"
)

func JiraTracker() tracker.LookupTracker {
	return tracker.Jira{
		tracker.Tracker{
			"Jira",
			"http://jira.com",
			[]*regexp.Regexp{
				regexp.MustCompile(`(EFFIG-[0-9])\s*`),
			},
		},
	}
}

func RallyTracker() tracker.LookupTracker {
	return tracker.Rally{
		tracker.Tracker{
			"Rally",
			"http://rally.com",
			[]*regexp.Regexp{
				regexp.MustCompile(`(US[0-9])\s*`),
				regexp.MustCompile(`(DE[0-9])\s*`),
			},
		},
	}
}

func TestTrackerGivenReturnsProperTrackerForStorySummaryLookup(t *testing.T) {
	trackers := []tracker.LookupTracker{JiraTracker(), RallyTracker()}

	assert.Equal(t, tracker.TrackerGiven("EFFIG-401", trackers).Details().Name, "Jira")
	assert.Equal(t, tracker.TrackerGiven("US123", trackers).Details().Name, "Rally")
	assert.Equal(t, tracker.TrackerGiven("DE123", trackers).Details().Name, "Rally")
}
