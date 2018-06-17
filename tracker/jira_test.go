package tracker_test

import (
	"regexp"
	"testing"

	"github.com/rafasf/story-summary/tracker"
	"github.com/stretchr/testify/assert"
)

func TheTracker() tracker.LookupTracker {
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
func TestStoryForReturnsStoryGivenIdentifier(t *testing.T) {
	jira := TheTracker()

	story := jira.StoryFor("EFFIG-401")
	expectedStory := tracker.Story{
		"EFFIG-401",
		"The Cool Summary",
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
	assert.Equal(t, "http://jira.com", jira.BaseUrl)
	assert.Equal(t, "(EFFIG-[0-9])\\s*", jira.Patterns[0].String())
}
