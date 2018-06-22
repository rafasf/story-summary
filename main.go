package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rafasf/story-summary/commit"
	"github.com/rafasf/story-summary/summary"
	"github.com/rafasf/story-summary/tracker"
)

func main() {
	logEntries := []string{
		"Add other to d|Rafael Ferreira|4ddf81f",
		"Add new to d|Rafael Ferreira|a319639",
		"EFFIG-401 First from other tracker|Rafael Ferreira|f861b45",
		"US11713 Aenean vestibulum ipsum et|Rafael Ferreira|cbe86a1",
		"US13791 Suspendisse dignissim hendrerit porttitor|Rafael Ferreira|cd4730f",
		"chore: Initial commit|Rafael Ferreira|a041e85"}

	tags := []commit.Tag{
		commit.Tag{regexp.MustCompile(`(US[0-9]+)\s*`), "Story"},
		commit.Tag{regexp.MustCompile(`(EFFIG-[0-9]+)\s*`), "Story"},
		commit.Tag{regexp.MustCompile(`chore:\s*`), "Chore"},
	}

	trackers := []tracker.LookupTracker{
		tracker.Jira{
			Info: tracker.Tracker{
				Name:    "Jira",
				BaseURL: "http://jira.com",
				Patterns: []*regexp.Regexp{
					regexp.MustCompile(`(EFFIG-[0-9])\s*`),
				},
			},
		},
		tracker.Rally{
			Info: tracker.Tracker{
				Name:    "Rally",
				BaseURL: "http://rally.com",
				Patterns: []*regexp.Regexp{
					regexp.MustCompile(`(US[0-9])\s*`),
					regexp.MustCompile(`(DE[0-9])\s*`),
				},
			},
		},
	}

	storyPatterns := allStoryPatternsGiven(trackers)

	isStory := func(s string) bool {
		return storyPatterns.MatchString(s)
	}

	allCommits := commit.CommitsFrom(logEntries, tags, "|")
	allCommitsByTag := commit.ByTag(allCommits)

	storyIds, generalCommits := summary.StoryIdsAndCommitsFrom(allCommitsByTag, isStory)

	fmt.Println(summary.For(storyIds, trackers))
	fmt.Println("==========")
	fmt.Println(generalCommits)
}

func allStoryPatternsGiven(trackers LookupTracker) *regexp.Regexp {
	var p []string
	for _, t := range trackers {
		p = append(p, t.AllPatterns().String())
	}
	storyPatterns, _ := regexp.Compile(strings.Join(p, "|"))

	return storyPatterns
}
