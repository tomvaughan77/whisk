package schedule

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	testCases := []struct {
		name               string
		inputCalendar      Year
		inputAvailableDays int
		expected           []Booking
		assertion          bool
	}{
		{
			name: "No leave is taken when no days are avaliable",
			inputCalendar: Year{
				days: []Day{
					{date: time.Now()}, // Monday Week 1
					{date: time.Now().AddDate(0, 0, 1)},
					{date: time.Now().AddDate(0, 0, 2)},
					{date: time.Now().AddDate(0, 0, 3)},
					{date: time.Now().AddDate(0, 0, 4)},
					{date: time.Now().AddDate(0, 0, 5), busy: true},
					{date: time.Now().AddDate(0, 0, 6), busy: true},
					{date: time.Now().AddDate(0, 0, 7)}, // Monday Week 2
					{date: time.Now().AddDate(0, 0, 8)},
					{date: time.Now().AddDate(0, 0, 9)},
					{date: time.Now().AddDate(0, 0, 10)},
					{date: time.Now().AddDate(0, 0, 11)},
					{date: time.Now().AddDate(0, 0, 12), busy: true},
					{date: time.Now().AddDate(0, 0, 13), busy: true},
				},
			},
			inputAvailableDays: 0,
			expected:           []Booking{},
			assertion:          true,
		},
		{
			name: "Last 2 days are taken in two week block when 2 avaliable",
			inputCalendar: Year{
				days: []Day{
					{date: time.Now()}, // Monday Week 1
					{date: time.Now().AddDate(0, 0, 1)},
					{date: time.Now().AddDate(0, 0, 2)},
					{date: time.Now().AddDate(0, 0, 3)},
					{date: time.Now().AddDate(0, 0, 4)},
					{date: time.Now().AddDate(0, 0, 5), busy: true},
					{date: time.Now().AddDate(0, 0, 6), busy: true},
					{date: time.Now().AddDate(0, 0, 7)}, // Monday Week 2
					{date: time.Now().AddDate(0, 0, 8)},
					{date: time.Now().AddDate(0, 0, 9)},
					{date: time.Now().AddDate(0, 0, 10)},
					{date: time.Now().AddDate(0, 0, 11)},
					{date: time.Now().AddDate(0, 0, 12), busy: true},
					{date: time.Now().AddDate(0, 0, 13), busy: true},
				},
			},
			inputAvailableDays: 2,
			expected: []Booking{
				{
					startDate:   time.Now().AddDate(0, 0, 10),
					endDate:     time.Now().AddDate(0, 0, 11),
					daysTaken:   2,
					breakLength: 4,
				},
			},
			assertion: true,
		},
		{
			name: "Second week in two week block when 5 avaliable",
			inputCalendar: Year{
				days: []Day{
					{date: time.Now()}, // Monday Week 1
					{date: time.Now().AddDate(0, 0, 1)},
					{date: time.Now().AddDate(0, 0, 2)},
					{date: time.Now().AddDate(0, 0, 3)},
					{date: time.Now().AddDate(0, 0, 4)},
					{date: time.Now().AddDate(0, 0, 5), busy: true},
					{date: time.Now().AddDate(0, 0, 6), busy: true},
					{date: time.Now().AddDate(0, 0, 7)}, // Monday Week 2
					{date: time.Now().AddDate(0, 0, 8)},
					{date: time.Now().AddDate(0, 0, 9)},
					{date: time.Now().AddDate(0, 0, 10)},
					{date: time.Now().AddDate(0, 0, 11)},
					{date: time.Now().AddDate(0, 0, 12), busy: true},
					{date: time.Now().AddDate(0, 0, 13), busy: true},
				},
			},
			inputAvailableDays: 2,
			expected: []Booking{
				{
					startDate:   time.Now().AddDate(0, 0, 7),
					endDate:     time.Now().AddDate(0, 0, 11),
					daysTaken:   5,
					breakLength: 9,
				},
			},
			assertion: true,
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Running test %d: %s", i, testCase.name)
		sch, err := Schedule(testCase.inputCalendar, testCase.inputAvailableDays)
		if err != nil {
			t.Errorf(err.Error())
		}

		assert.Equal(t, testCase.expected, reflect.DeepEqual(testCase.assertion, sch), fmt.Sprintf("Test %d failed", i))
	}
}
