package scheduler

import (
	"testing"
	"time"
)

func TestGetDurationAndGetCron(t *testing.T) {
	timeRange1 := TimeRange{
		HourMinuteFrom: "09:00",
		HourMinuteTo:   "03:00",
		Frequency:      DAILY,
	}
	wantDuration1 := 18 * time.Hour
	wantCron1 := "00 09 * * *"
	timeRange2 := TimeRange{
		HourMinuteFrom: "12:00",
		HourMinuteTo:   "14:00",
		Frequency:      DAILY,
	}
	wantDuration2 := 2 * time.Hour
	wantCron2 := "00 12 * * *"

	timeRange3 := TimeRange{
		HourMinuteFrom: "14:00",
		HourMinuteTo:   "12:00",
		Frequency:      DAILY,
	}
	wantDuration3 := 22 * time.Hour
	wantCron3 := "00 14 * * *"
	timeRange4 := TimeRange{
		HourMinuteFrom: "09:00",
		HourMinuteTo:   "12:00", //00:01
		Weekdays:       []time.Weekday{time.Monday, time.Tuesday, time.Wednesday},
		Frequency:      WEEKLY,
	}
	wantDuration4 := 3 * time.Hour
	wantCron4 := "00 09 * * 1,2,3"
	timeRange5 := TimeRange{
		HourMinuteFrom: "09:00",
		HourMinuteTo:   "00:00",
		Frequency:      WEEKLY,
		Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
	}
	wantDuration5 := 15 * time.Hour
	wantCron5 := "00 09 * * 1,3,5"
	timeRange6 := TimeRange{
		HourMinuteFrom: "17:00",
		HourMinuteTo:   "19:30",
		Frequency:      WEEKLY,
		Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
	}
	wantDuration6 := 2*time.Hour + 30*time.Minute
	wantCron6 := "00 17 * * 1,3,5"
	timeRange7 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "01:00",
		Frequency:      WEEKLY_RANGE,
		WeekdayFrom:    4,
		WeekdayTo:      1,
	}
	wantDuration7 := 86 * time.Hour
	wantCron7 := "00 11 * * 4"
	timeRange8 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "13:00",
		Frequency:      WEEKLY_RANGE,
		WeekdayFrom:    1,
		WeekdayTo:      4,
	}
	wantDuration8 := 74 * time.Hour
	wantCron8 := "00 11 * * 1"
	timeRange9 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "12:00",
		Frequency:      WEEKLY_RANGE,
		WeekdayTo:      9,
		WeekdayFrom:    1,
	}
	wantDuration9 := 0 * time.Hour
	wantCron9 := "00 11 * * 1"
	timeRange10 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "12:00",
		Frequency:      WEEKLY_RANGE,
		WeekdayTo:      1,
		WeekdayFrom:    9,
	}
	wantDuration10 := 0 * time.Second
	wantCron10 := "00 11 * * 9"
	timeRange11 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "12:00",
		DayFrom:        1,
		DayTo:          3,
		Frequency:      MONTHLY,
	}
	wantDuration11 := 49 * time.Hour
	wantCron11 := "00 11 1 * *"
	timeRange12 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "00:00",
		DayFrom:        1,
		DayTo:          3,
		Frequency:      MONTHLY,
	}
	wantDuration12 := 37 * time.Hour
	wantCron12 := "00 11 1 * *"
	timeRange13 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "02:00",
		DayFrom:        1,
		DayTo:          3,
		Frequency:      MONTHLY,
	}
	wantDuration13 := 39 * time.Hour
	wantCron13 := "00 11 1 * *"
	timeRange14 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "02:00",
		DayFrom:        27,
		DayTo:          -2,
		Frequency:      MONTHLY,
	}
	wantDuration14 := 63 * time.Hour
	wantCron14 := "00 11 27 * *"
	timeRange15 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "02:00",
		DayFrom:        27,
		DayTo:          -4,
		Frequency:      MONTHLY,
	}
	wantDuration15 := 15 * time.Hour
	wantCron15 := "00 11 27 * *"
	timeRange16 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "12:00",
		DayFrom:        27,
		DayTo:          -4,
		Frequency:      MONTHLY,
	}
	wantDuration16 := 25 * time.Hour
	wantCron16 := "00 11 27 * *"
	timeRange17 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "11:00",
		DayFrom:        27,
		DayTo:          -4,
		Frequency:      MONTHLY,
	}
	wantDuration17 := 24 * time.Hour
	wantCron17 := "00 11 27 * *"
	timeRange18 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "13:00",
		DayFrom:        27,
		DayTo:          4,
		Frequency:      MONTHLY,
	}
	wantDuration18 := 194 * time.Hour
	wantCron18 := "00 11 27 * *"
	timeRange19 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "12:00",
		DayFrom:        27,
		DayTo:          4,
		Frequency:      MONTHLY,
	}
	wantDuration19 := 193 * time.Hour
	wantCron19 := "00 11 27 * *"
	timeRange20 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "13:00",
		DayFrom:        -2,
		DayTo:          -4,
		Frequency:      MONTHLY,
	}
	wantDuration20 := 698 * time.Hour
	wantCron20 := "00 11 L-1 * *"
	timeRange21 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "13:00",
		DayFrom:        -2,
		DayTo:          -1,
		Frequency:      MONTHLY,
	}
	wantDuration21 := 26 * time.Hour
	wantCron21 := "00 11 L-1 * *"
	timeRange22 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "13:00",
		DayFrom:        -2,
		DayTo:          4,
		Frequency:      MONTHLY,
	}
	wantDuration22 := 122 * time.Hour
	wantCron22 := "00 11 L-1 * *"
	timeRange23 := TimeRange{
		HourMinuteFrom: "09:00",
		HourMinuteTo:   "00:01", //00:01
		Weekdays:       []time.Weekday{time.Monday, time.Tuesday, time.Wednesday},
		Frequency:      WEEKLY,
	}
	wantDuration23 := 15*time.Hour + 1*time.Minute
	wantCron23 := "00 09 * * 1,2,3"
	timeRange24 := TimeRange{
		HourMinuteFrom: "12:00ab",
		HourMinuteTo:   "14:00ab",
		Frequency:      DAILY,
	}
	wantDuration24 := 0 * time.Hour
	wantCron24 := "00ab 12 * * *"
	tests := []struct {
		timeRange    TimeRange
		wantDuration time.Duration
		wantCron     string
	}{
		{timeRange1, wantDuration1, wantCron1},
		{timeRange2, wantDuration2, wantCron2},
		{timeRange3, wantDuration3, wantCron3},
		{timeRange4, wantDuration4, wantCron4},
		{timeRange5, wantDuration5, wantCron5},
		{timeRange6, wantDuration6, wantCron6},
		{timeRange7, wantDuration7, wantCron7},
		{timeRange8, wantDuration8, wantCron8},
		{timeRange9, wantDuration9, wantCron9},
		{timeRange10, wantDuration10, wantCron10},
		{timeRange11, wantDuration11, wantCron11},
		{timeRange12, wantDuration12, wantCron12},
		{timeRange13, wantDuration13, wantCron13},
		{timeRange14, wantDuration14, wantCron14},
		{timeRange15, wantDuration15, wantCron15},
		{timeRange16, wantDuration16, wantCron16},
		{timeRange17, wantDuration17, wantCron17},
		{timeRange18, wantDuration18, wantCron18},
		{timeRange19, wantDuration19, wantCron19},
		{timeRange20, wantDuration20, wantCron20},
		{timeRange21, wantDuration21, wantCron21},
		{timeRange22, wantDuration22, wantCron22},
		{timeRange23, wantDuration23, wantCron23},
		{timeRange24, wantDuration24, wantCron24},
	}

	for i, test := range tests {
		// Test getDuration method
		gotDuration, err := test.timeRange.getDuration(12, 2024)
		if err != nil {
			if gotDuration != test.wantDuration {
				t.Errorf("Test case %d: getDuration() = %v, want %v", i+1, gotDuration, test.wantDuration)
			}
		}
		if gotDuration != test.wantDuration {
			t.Errorf("Test case %d: getDuration() = %v, want %v", i+1, gotDuration, test.wantDuration)
		}

		// Test getCron method
		gotCron := test.timeRange.getCron()
		if gotCron != test.wantCron {
			t.Errorf("Test case %d: getCron() = %v, want %v", i+1, gotCron, test.wantCron)
		}

	}

}
