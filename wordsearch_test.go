package main

import (
	"fmt"
	"testing"
)

var sut = Graph{
	{"H", "B", "C", "D", "D"},
	{"D", "E", "F", "G", "H"},
	{"G", "H", "L", "M", "N"},
	{"H", "B", "C", "L", "L"},
	{"D", "E", "F", "O", "O"},
}

func TestNorthGetNextPosition(t *testing.T) {
	type testCase struct {
		cursorX int
		expectedInbounds bool
		expectedX int
	}

	var testCases = []testCase {
		{
			cursorX: -5,
			expectedInbounds: false,
			expectedX: -5,
		},
		{
			cursorX: -2,
			expectedInbounds: false,
			expectedX: -2,
		},
		{
			cursorX: -1,
			expectedInbounds: true,
			expectedX: 0,
		},
		{
			cursorX: 0,
			expectedInbounds: true,
			expectedX: 1,
		},
		{
			cursorX: 1,
			expectedInbounds: true,
			expectedX: 2,
		},
		{
			cursorX: 3,
			expectedInbounds: true,
			expectedX: 4,
		},
		{
			cursorX: 4,
			expectedInbounds: false,
			expectedX: 4,
		},
		{
			cursorX: 5,
			expectedInbounds: false,
			expectedX: 5,
		},
		{
			cursorX: 20,
			expectedInbounds: false,
			expectedX: 20,
		},
	}

	const cursorY = 0
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("(%d,%d)", testCase.cursorX, cursorY), func(t *testing.T) {
			t.Parallel()
			inBounds, cursorX, _ := sut.GetNextPosition(testCase.cursorX, cursorY, North)
			if inBounds != testCase.expectedInbounds {
				t.Fatalf("Expected inBounds to be %t, actual was %t", testCase.expectedInbounds, inBounds)
			}
			if cursorX != testCase.expectedX {
				t.Fatalf("Expected cursorX to be %d, actual was %d", testCase.expectedX, cursorX)
			}
		})
	}
}