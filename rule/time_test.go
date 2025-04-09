package rule

import (
	"testing"
	"time"
)

func TestTimeBetween(t *testing.T) {
	now := time.Now()
	before := now.Add(-time.Hour)
	after := now.Add(time.Hour)

	tests := []struct {
		name    string
		rule    *TimeBetweenRule
		value   time.Time
		wantErr bool
	}{
		{
			name:    "valid: time is between",
			rule:    TimeBetween(before, after),
			value:   now,
			wantErr: false,
		},
		{
			name:    "invalid: time is before start",
			rule:    TimeBetween(now, after),
			value:   before,
			wantErr: true,
		},
		{
			name:    "invalid: time is after end",
			rule:    TimeBetween(before, now),
			value:   after,
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    TimeBetween(before, after).Errf("custom error"),
			value:   before.Add(-time.Hour),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeBetweenRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBefore(t *testing.T) {
	now := time.Now()
	before := now.Add(-time.Hour)
	after := now.Add(time.Hour)

	tests := []struct {
		name    string
		rule    *BeforeRule
		value   time.Time
		wantErr bool
	}{
		{
			name:    "valid: time is before",
			rule:    Before(now),
			value:   before,
			wantErr: false,
		},
		{
			name:    "invalid: time is after",
			rule:    Before(now),
			value:   after,
			wantErr: true,
		},
		{
			name:    "invalid: time is equal",
			rule:    Before(now),
			value:   now,
			wantErr: true,
		},
		{
			name:    "valid: time is equal with IncludeTime",
			rule:    Before(now).IncludeTime(),
			value:   now,
			wantErr: false,
		},
		{
			name:    "custom error message",
			rule:    Before(now).Errf("custom error"),
			value:   after,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeforeRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAfter(t *testing.T) {
	now := time.Now()
	before := now.Add(-time.Hour)
	after := now.Add(time.Hour)

	tests := []struct {
		name    string
		rule    *AfterRule
		value   time.Time
		wantErr bool
	}{
		{
			name:    "valid: time is after",
			rule:    After(now),
			value:   after,
			wantErr: false,
		},
		{
			name:    "invalid: time is before",
			rule:    After(now),
			value:   before,
			wantErr: true,
		},
		{
			name:    "invalid: time is equal",
			rule:    After(now),
			value:   now,
			wantErr: true,
		},
		{
			name:    "valid: time is equal with IncludeTime",
			rule:    After(now).IncludeTime(),
			value:   now,
			wantErr: false,
		},
		{
			name:    "custom error message",
			rule:    After(now).Errf("custom error"),
			value:   before,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AfterRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDateFormat(t *testing.T) {
	tests := []struct {
		name    string
		rule    *DateFormatRule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: YYYY-MM-DD",
			rule:    DateFormat("2006-01-02"),
			value:   "2024-03-15",
			wantErr: false,
		},
		{
			name:    "valid: DD/MM/YYYY",
			rule:    DateFormat("02/01/2006"),
			value:   "15/03/2024",
			wantErr: false,
		},
		{
			name:    "invalid: wrong format",
			rule:    DateFormat("2006-01-02"),
			value:   "2024/03/15",
			wantErr: true,
		},
		{
			name:    "invalid: invalid date",
			rule:    DateFormat("2006-01-02"),
			value:   "2024-13-45",
			wantErr: true,
		},
		{
			name:    "valid: empty string",
			rule:    DateFormat("2006-01-02"),
			value:   "",
			wantErr: false,
		},
		{
			name:    "custom error message",
			rule:    DateFormat("2006-01-02").Errf("custom error"),
			value:   "2024/03/15",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateFormatRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimeFormat(t *testing.T) {
	tests := []struct {
		name    string
		rule    *TimeFormatRule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: HH:mm:ss",
			rule:    TimeFormat("15:04:05"),
			value:   "14:30:00",
			wantErr: false,
		},
		{
			name:    "valid: HH:mm",
			rule:    TimeFormat("15:04"),
			value:   "14:30",
			wantErr: false,
		},
		{
			name:    "invalid: wrong format",
			rule:    TimeFormat("15:04:05"),
			value:   "14:30",
			wantErr: true,
		},
		{
			name:    "invalid: invalid time",
			rule:    TimeFormat("15:04:05"),
			value:   "25:61:99",
			wantErr: true,
		},
		{
			name:    "valid: empty string",
			rule:    TimeFormat("15:04:05"),
			value:   "",
			wantErr: false,
		},
		{
			name:    "custom error message",
			rule:    TimeFormat("15:04:05").Errf("custom error"),
			value:   "14:30",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeFormatRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDateTimeFormat(t *testing.T) {
	tests := []struct {
		name    string
		rule    *DateTimeFormatRule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: YYYY-MM-DD HH:mm:ss",
			rule:    DateTimeFormat("2006-01-02 15:04:05"),
			value:   "2024-03-15 14:30:00",
			wantErr: false,
		},
		{
			name:    "valid: DD/MM/YYYY HH:mm",
			rule:    DateTimeFormat("02/01/2006 15:04"),
			value:   "15/03/2024 14:30",
			wantErr: false,
		},
		{
			name:    "invalid: wrong format",
			rule:    DateTimeFormat("2006-01-02 15:04:05"),
			value:   "2024/03/15 14:30:00",
			wantErr: true,
		},
		{
			name:    "invalid: invalid datetime",
			rule:    DateTimeFormat("2006-01-02 15:04:05"),
			value:   "2024-13-45 25:61:99",
			wantErr: true,
		},
		{
			name:    "valid: empty string",
			rule:    DateTimeFormat("2006-01-02 15:04:05"),
			value:   "",
			wantErr: false,
		},
		{
			name:    "custom error message",
			rule:    DateTimeFormat("2006-01-02 15:04:05").Errf("custom error"),
			value:   "2024/03/15 14:30:00",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTimeFormatRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeekend(t *testing.T) {
	// a saturday date
	saturday := time.Date(2024, 3, 16, 0, 0, 0, 0, time.Local)
	// a sunday date
	sunday := time.Date(2024, 3, 17, 0, 0, 0, 0, time.Local)
	// a workday date
	workday := time.Date(2024, 3, 15, 0, 0, 0, 0, time.Local)

	tests := []struct {
		name    string
		rule    *WeekendRule
		value   time.Time
		wantErr bool
	}{
		{
			name:    "valid: saturday",
			rule:    Weekend(),
			value:   saturday,
			wantErr: false,
		},
		{
			name:    "valid: sunday",
			rule:    Weekend(),
			value:   sunday,
			wantErr: false,
		},
		{
			name:    "invalid: workday",
			rule:    Weekend(),
			value:   workday,
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    Weekend().Errf("custom error"),
			value:   workday,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeekendRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWorkday(t *testing.T) {
	// a saturday date
	saturday := time.Date(2024, 3, 16, 0, 0, 0, 0, time.Local)
	// a sunday date
	sunday := time.Date(2024, 3, 17, 0, 0, 0, 0, time.Local)
	// a workday date
	workday := time.Date(2024, 3, 15, 0, 0, 0, 0, time.Local)

	tests := []struct {
		name    string
		rule    *WorkdayRule
		value   time.Time
		wantErr bool
	}{
		{
			name:    "valid: workday",
			rule:    Workday(),
			value:   workday,
			wantErr: false,
		},
		{
			name:    "invalid: saturday",
			rule:    Workday(),
			value:   saturday,
			wantErr: true,
		},
		{
			name:    "invalid: sunday",
			rule:    Workday(),
			value:   sunday,
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    Workday().Errf("custom error"),
			value:   saturday,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("WorkdayRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHoliday(t *testing.T) {
	// a holiday date
	holiday := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	// a workday date
	workday := time.Date(2024, 3, 15, 0, 0, 0, 0, time.Local)
	// a weekend date
	weekend := time.Date(2024, 3, 16, 0, 0, 0, 0, time.Local)

	tests := []struct {
		name    string
		rule    *HolidayRule
		value   time.Time
		wantErr bool
	}{
		{
			name:    "valid: holiday",
			rule:    Holiday(holiday),
			value:   holiday,
			wantErr: false,
		},
		{
			name:    "valid: weekend",
			rule:    Holiday(holiday),
			value:   weekend,
			wantErr: false,
		},
		{
			name:    "invalid: workday",
			rule:    Holiday(holiday),
			value:   workday,
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    Holiday(holiday).Errf("custom error"),
			value:   workday,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HolidayRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
