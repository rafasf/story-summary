package summary_test

import (
	"regexp"
	"testing"

	"github.com/rafasf/story-summary/commit"
	"github.com/rafasf/story-summary/summary"
	"github.com/rafasf/story-summary/tracker"
	"github.com/stretchr/testify/assert"
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

func TestStoryIdsAndCommitsFromReturnsListsWithRespectiveInfo(t *testing.T) {
	storyTag := commit.TagMatch{
		"A123",
		commit.Tag{regexp.MustCompile(`(A[0-9])\s*`), "Story"}}
	generalTag := commit.TagMatch{
		"General",
		commit.Tag{regexp.MustCompile(`^General$`), "General"}}

	isStory := func(s string) bool {
		return storyTag.Pattern().MatchString(s)
	}

	generalCommit := commit.Commit{"Subject 2", "Bob", "hash2", generalTag}
	someCommitsByTag := map[string][]commit.Commit{
		"A123": []commit.Commit{
			commit.Commit{"Subject 1", "Bob", "hash1", storyTag},
			commit.Commit{"Subject 1", "Bob", "hash1", storyTag},
		},
		"General": []commit.Commit{generalCommit},
	}

	storyIds, generalCommits := summary.StoryIdsAndCommitsFrom(
		someCommitsByTag,
		isStory)

	assert.Equal(t, []string{"A123"}, storyIds)
	assert.Equal(t, []commit.Commit{generalCommit}, generalCommits)
}
