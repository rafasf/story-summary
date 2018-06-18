package main

import (
	"fmt"
	"regexp"

	"github.com/rafasf/story-summary/commit"
	//	"github.com/rafasf/story-summary/tracker"
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

	// trackers := []tracker.LookupTracker{
	// 	tracker.Jira{
	// 		tracker.Tracker{
	// 			"Jira",
	// 			"http://jira.com",
	// 			[]*regexp.Regexp{
	// 				regexp.MustCompile(`(EFFIG-[0-9])\s*`),
	// 			},
	// 		},
	// 	},
	// 	tracker.Rally{
	// 		tracker.Tracker{
	// 			"Rally",
	// 			"http://rally.com",
	// 			[]*regexp.Regexp{
	// 				regexp.MustCompile(`(US[0-9])\s*`),
	// 				regexp.MustCompile(`(DE[0-9])\s*`),
	// 			},
	// 		},
	// 	},
	// }

	commits := commit.CommitsFrom(logEntries, tags, "|")

	for _, commit := range commits {
		fmt.Println(commit)
	}

	//fmt.Println(ByTag(commits).Keys())
}
