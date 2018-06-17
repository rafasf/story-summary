package tracker_test

import (
	"regexp"
	"testing"

	"github.com/rafasf/story-summary/tracker"
	"github.com/stretchr/testify/assert"
)

func TheTracker() tracker.LookupTracker {
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
func TestStoryForReturnsStoryGivenIdentifier(t *testing.T) {
	jira := TheTracker()

	story := jira.StoryFor("EFFIG-401")
	expectedStory := tracker.Story{
		Identifier: "EFFIG-401",
		Summary:    "The Cool Summary",
	}

	assert.Equal(t, expectedStory, story)
}

func TestAllPatternsReturnsCombinedRegex(t *testing.T) {
	jira := TheTracker()

	assert.Equal(t, "(EFFIG-[0-9])\\s*", jira.AllPatterns().String())
}

func TestDetailsReturnsJiraInformation(t *testing.T) {
	jira := TheTracker().Details()

	assert.Equal(t, "Jira", jira.Name)
	assert.Equal(t, "http://jira.com", jira.BaseURL)
	assert.Equal(t, "(EFFIG-[0-9])\\s*", jira.Patterns[0].String())
}
