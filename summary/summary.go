package summary

import (
	"github.com/rafasf/story-summary/tracker"
)

// StorySummary represent the unique stories present in the changelog
// along with its summaries from the tracker.
type StorySummary struct {
	Stories []tracker.Story
}

// For returns a StorySummary, fetching the summary for each of the stories
// from the proper tracker.
func For(storyIds []string, trackers []tracker.LookupTracker) StorySummary {
	var stories []tracker.Story
	for _, storyID := range storyIds {
		possibleTracker := tracker.Given(storyID, trackers)

		if possibleTracker != nil {
			stories = append(stories, possibleTracker.StoryFor(storyID))
		} else {
			stories = append(stories, tracker.Story{Identifier: storyID})
		}
	}
	return StorySummary{stories}
}
