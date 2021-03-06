// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import regexp "regexp"
import tracker "github.com/rafasf/story-summary/tracker"

// LookupTracker is an autogenerated mock type for the LookupTracker type
type LookupTracker struct {
	mock.Mock
}

// AllPatterns provides a mock function with given fields:
func (_m *LookupTracker) AllPatterns() *regexp.Regexp {
	ret := _m.Called()

	var r0 *regexp.Regexp
	if rf, ok := ret.Get(0).(func() *regexp.Regexp); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*regexp.Regexp)
		}
	}

	return r0
}

// Details provides a mock function with given fields:
func (_m *LookupTracker) Details() tracker.Tracker {
	ret := _m.Called()

	var r0 tracker.Tracker
	if rf, ok := ret.Get(0).(func() tracker.Tracker); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(tracker.Tracker)
	}

	return r0
}

// StoryFor provides a mock function with given fields: storyID
func (_m *LookupTracker) StoryFor(storyID string) tracker.Story {
	ret := _m.Called(storyID)

	var r0 tracker.Story
	if rf, ok := ret.Get(0).(func(string) tracker.Story); ok {
		r0 = rf(storyID)
	} else {
		r0 = ret.Get(0).(tracker.Story)
	}

	return r0
}
