package summary_test

import (
	"github.com/rafasf/story-summary/summary"
	"github.com/rafasf/story-summary/tracker"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func JiraTracker() tracker.LookupTracker {
	return tracker.Jira{
		Info: tracker.Tracker{
			Name:    "A Tracker",
			BaseURL: "http://a-tracker.fake",
			Patterns: []*regexp.Regexp{
				regexp.MustCompile(`(A[0-9])\s*`),
			},
		},
	}
}

func TestForReturnsStorySummariesForIdentifiers(t *testing.T) {
	trackers := []tracker.LookupTracker{JiraTracker()}
	identifiers := []string{"A123"}

	storySummaries := summary.For(identifiers, trackers)
	expectedSummaries := summary.StorySummary{
		Stories: []tracker.Story{
			tracker.Story{Summary: "The Cool Summary", Identifier: "A123"},
		},
	}

	assert.Equal(t, expectedSummaries, storySummaries)
}
