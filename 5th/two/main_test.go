package main

import "testing"

func TestSeatID(t *testing.T) {
	testCases := map[string]struct {
		seat   string
		seatID int64
	}{
		"eg 1": {
			seat:   "FBFBBFFRLR",
			seatID: 357,
		},
		"eg 2": {
			seat:   "BFFFBBFRRR",
			seatID: 567,
		},
		"eg 3": {
			seat:   "FFFBBBFRRR",
			seatID: 119,
		},
		"eg 4": {
			seat:   "BBFFBBFRLL",
			seatID: 820,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			seatID := seatID(tc.seat)
			if seatID != tc.seatID {
				t.Errorf("unexpected seat id '%d', expected '%d'", seatID, tc.seatID)
			}
		})
	}
}
