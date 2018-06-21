package summary

import (
	"github.com/rafasf/story-summary/commit"
	"github.com/rafasf/story-summary/tracker"
)

// StorySummary represent the unique stories present in the changelog
// along with its summaries from the tracker.
type StorySummary struct {
	Stories []tracker.Story
}

// StoryIdsAndCommitsFrom returns a unique list of story identifiers and
// a list of commits that don't belong to any story.
func StoryIdsAndCommitsFrom(
	commitsByTag map[string][]commit.Commit,
	isStory func(string) bool,
) ([]string, []commit.Commit) {
	var storyIds []string
	var generalCommits []commit.Commit

	for tag, commits := range commitsByTag {
		if isStory(tag) {
			storyIds = append(storyIds, tag)
		} else {
			generalCommits = append(generalCommits, commits...)
		}
	}

	return storyIds, generalCommits
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
