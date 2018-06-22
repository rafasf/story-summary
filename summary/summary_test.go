package summary_test

import (
	"regexp"
	"testing"

	"github.com/rafasf/story-summary/commit"
	"github.com/rafasf/story-summary/mocks"
	"github.com/rafasf/story-summary/summary"
	"github.com/rafasf/story-summary/tracker"
	"github.com/stretchr/testify/assert"
)

func TestForReturnsStorySummariesForIdentifiers(t *testing.T) {
	aTracker := new(mocks.LookupTracker)
	trackers := []tracker.LookupTracker{aTracker}
	identifiers := []string{"A123"}

	story := tracker.Story{Summary: "The Cool Summary", Identifier: "A123"}
	expectedSummary := summary.StorySummary{
		Stories: []tracker.Story{story},
	}

	aTracker.On("AllPatterns").Return(regexp.MustCompile(`(A[0-9])\s*`))
	aTracker.On("StoryFor", "A123").Return(story)

	storySummaries := summary.For(identifiers, trackers)

	assert.Equal(t, expectedSummary, storySummaries)
	aTracker.AssertExpectations(t)
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
