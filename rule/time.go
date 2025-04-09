// Package rule provides a collection of validation rules for various data types.
// This file contains time-related validation rules for dates, times, and time ranges.
package rule

import (
	"errors"
	"fmt"
	"time"
)

// Time validation errors
var (
	// ErrTimeBetween is returned when a time value is outside a specified range.
	// The time must be after the start time and before the end time.
	ErrTimeBetween = errors.New("time must be between the specified times")

	// ErrBefore is returned when a time value is not before a specified time.
	// By default, the time must be strictly before the reference time.
	ErrBefore = errors.New("time must be before the specified time")

	// ErrAfter is returned when a time value is not after a specified time.
	// By default, the time must be strictly after the reference time.
	ErrAfter = errors.New("time must be after the specified time")

	// ErrDateFormat is returned when a string does not match the expected date format.
	// The format should follow Go's time format specification.
	ErrDateFormat = errors.New("invalid date format")

	// ErrTimeFormat is returned when a string does not match the expected time format.
	// The format should follow Go's time format specification.
	ErrTimeFormat = errors.New("invalid time format")

	// ErrDateTimeFormat is returned when a string does not match the expected datetime format.
	// The format should follow Go's time format specification.
	ErrDateTimeFormat = errors.New("invalid datetime format")

	// ErrWeekend is returned when a time value is not a weekend day (Saturday or Sunday).
	ErrWeekend = errors.New("time must be a weekend")

	// ErrWorkday is returned when a time value is not a workday (Monday through Friday).
	ErrWorkday = errors.New("time must be a workday")

	// ErrHoliday is returned when a time value is not in the list of specified holidays.
	ErrHoliday = errors.New("time must be a holiday")
)

// TimeBetweenRule validates that a time falls within a specified range.
// The time must be after the start time and before the end time.
//
// Example:
//
//	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
//	rule := TimeBetween(start, end).Err("Date must be in 2023")
//	err := rule.Validate(time.Now())  // returns nil if current time is in 2023
type TimeBetweenRule struct {
	start time.Time
	end   time.Time
	e     error
}

// TimeBetween creates a new time range validation rule.
// The rule ensures that a time value falls between the specified start and end times.
//
// Example:
//
//	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
//	rule := TimeBetween(start, end)
func TimeBetween(start, end time.Time) *TimeBetweenRule {
	return &TimeBetweenRule{
		start: start,
		end:   end,
		e:     ErrTimeBetween,
	}
}

// Validate checks if the given time falls within the specified range.
// Returns nil if the time is between start and end (inclusive), or an error otherwise.
//
// Example:
//
//	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
//	rule := TimeBetween(start, end)
//	err := rule.Validate(time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC))  // returns nil
//	err = rule.Validate(time.Date(2022, 12, 31, 23, 59, 59, 0, time.UTC))  // returns error
func (r *TimeBetweenRule) Validate(value time.Time) error {
	if value.Before(r.start) || value.After(r.end) {
		if r.e != nil {
			return r.e
		}
		return ErrTimeBetween
	}
	return nil
}

// Errf sets a custom error message for time range validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
//	rule := TimeBetween(start, end).Errf("Event date must be in 2023")
func (r *TimeBetweenRule) Errf(format string, args ...any) *TimeBetweenRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// BeforeRule validates that a time is before a specified reference time.
// By default, the time must be strictly before the reference time.
//
// Example:
//
//	deadline := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
//	rule := Before(deadline).Errf("Submission must be before the deadline")
//	err := rule.Validate(time.Now())  // returns nil if current time is before deadline
type BeforeRule struct {
	t           time.Time
	includeTime bool
	e           error
}

// Before creates a new "before time" validation rule.
// The rule ensures that a time value is before the specified reference time.
//
// Example:
//
//	deadline := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
//	rule := Before(deadline)
func Before(t time.Time) *BeforeRule {
	return &BeforeRule{
		t:           t,
		includeTime: false,
		e:           ErrBefore,
	}
}

// IncludeTime modifies the rule to include the reference time in the valid range.
// By default, the time must be strictly before the reference time.
//
// Example:
//
//	deadline := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
//	rule := Before(deadline).IncludeTime()  // allows time to be equal to deadline
func (r *BeforeRule) IncludeTime() *BeforeRule {
	r.includeTime = true
	return r
}

// Validate checks if the given time is before the reference time.
// If IncludeTime() was called, the time can also be equal to the reference time.
//
// Example:
//
//	deadline := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
//	rule := Before(deadline)
//	err := rule.Validate(time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC))  // returns nil
//	err = rule.Validate(deadline)  // returns error
//
//	rule = Before(deadline).IncludeTime()
//	err = rule.Validate(deadline)  // returns nil
func (r *BeforeRule) Validate(value time.Time) error {
	if r.includeTime {
		if !value.Before(r.t) && !value.Equal(r.t) {
			if r.e != nil {
				return r.e
			}
			return ErrBefore
		}
	} else {
		if !value.Before(r.t) {
			if r.e != nil {
				return r.e
			}
			return ErrBefore
		}
	}
	return nil
}

// Errf sets a custom error message for "before time" validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	deadline := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
//	rule := Before(deadline).Errf("Submission must be before the deadline")
func (r *BeforeRule) Errf(format string, args ...any) *BeforeRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// AfterRule validates that a time is after a specified reference time.
// By default, the time must be strictly after the reference time.
//
// Example:
//
//	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	rule := After(startDate).Errf("Event must start after January 1, 2023")
//	err := rule.Validate(time.Now())  // returns nil if current time is after start date
type AfterRule struct {
	t           time.Time
	includeTime bool
	e           error
}

// After creates a new "after time" validation rule.
// The rule ensures that a time value is after the specified reference time.
//
// Example:
//
//	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	rule := After(startDate)
func After(t time.Time) *AfterRule {
	return &AfterRule{
		t:           t,
		includeTime: false,
		e:           ErrAfter,
	}
}

// IncludeTime modifies the rule to include the reference time in the valid range.
// By default, the time must be strictly after the reference time.
//
// Example:
//
//	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	rule := After(startDate).IncludeTime()  // allows time to be equal to start date
func (r *AfterRule) IncludeTime() *AfterRule {
	r.includeTime = true
	return r
}

// Validate checks if the given time is after the reference time.
// If IncludeTime() was called, the time can also be equal to the reference time.
//
// Example:
//
//	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	rule := After(startDate)
//	err := rule.Validate(time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC))  // returns nil
//	err = rule.Validate(startDate)  // returns error
//
//	rule = After(startDate).IncludeTime()
//	err = rule.Validate(startDate)  // returns nil
func (r *AfterRule) Validate(value time.Time) error {
	if r.includeTime {
		if !value.After(r.t) && !value.Equal(r.t) {
			if r.e != nil {
				return r.e
			}
			return ErrAfter
		}
	} else {
		if !value.After(r.t) {
			if r.e != nil {
				return r.e
			}
			return ErrAfter
		}
	}
	return nil
}

// Errf sets a custom error message for "after time" validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	rule := After(startDate).Errf("Event must start after January 1, 2023")
func (r *AfterRule) Errf(format string, args ...any) *AfterRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// DateFormatRule validates that a string matches a specified date format.
// The format should follow Go's time format specification.
//
// Example:
//
//	rule := DateFormat("2006-01-02").Errf("Date must be in YYYY-MM-DD format")
//	err := rule.Validate("2023-12-31")  // returns nil
//	err = rule.Validate("12/31/2023")  // returns error
type DateFormatRule struct {
	format string
	e      error
}

// DateFormat creates a new date format validation rule.
// The format should follow Go's time format specification (e.g., "2006-01-02").
//
// Example:
//
//	rule := DateFormat("2006-01-02")  // YYYY-MM-DD format
//	rule := DateFormat("02/01/2006")  // DD/MM/YYYY format
func DateFormat(format string) *DateFormatRule {
	return &DateFormatRule{
		format: format,
		e:      ErrDateFormat,
	}
}

// Validate checks if the given string matches the specified date format.
// Empty strings are considered valid.
//
// Example:
//
//	rule := DateFormat("2006-01-02")
//	err := rule.Validate("2023-12-31")  // returns nil
//	err = rule.Validate("12/31/2023")  // returns error
//	err = rule.Validate("")  // returns nil (empty string is valid)
func (r *DateFormatRule) Validate(value string) error {
	if value == "" {
		return nil
	}
	_, err := time.Parse(r.format, value)
	if err != nil {
		if r.e != nil {
			return r.e
		}
		return ErrDateFormat
	}
	return nil
}

// Errf sets a custom error message for date format validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := DateFormat("2006-01-02").Errf("Please enter the date in YYYY-MM-DD format")
func (r *DateFormatRule) Errf(format string, args ...any) *DateFormatRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// TimeFormatRule validates that a string matches a specified time format.
// The format should follow Go's time format specification.
//
// Example:
//
//	rule := TimeFormat("15:04").Errf("Time must be in HH:MM format")
//	err := rule.Validate("14:30")  // returns nil
//	err = rule.Validate("2:30 PM")  // returns error
type TimeFormatRule struct {
	format string
	e      error
}

// TimeFormat creates a new time format validation rule.
// The format should follow Go's time format specification (e.g., "15:04").
//
// Example:
//
//	rule := TimeFormat("15:04")  // 24-hour format (HH:MM)
//	rule := TimeFormat("03:04 PM")  // 12-hour format with AM/PM
func TimeFormat(format string) *TimeFormatRule {
	return &TimeFormatRule{
		format: format,
		e:      ErrTimeFormat,
	}
}

// Validate checks if the given string matches the specified time format.
// Empty strings are considered valid.
//
// Example:
//
//	rule := TimeFormat("15:04")
//	err := rule.Validate("14:30")  // returns nil
//	err = rule.Validate("2:30 PM")  // returns error
//	err = rule.Validate("")  // returns nil (empty string is valid)
func (r *TimeFormatRule) Validate(value string) error {
	if value == "" {
		return nil
	}
	_, err := time.Parse(r.format, value)
	if err != nil {
		if r.e != nil {
			return r.e
		}
		return ErrTimeFormat
	}
	return nil
}

// Errf sets a custom error message for time format validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := TimeFormat("15:04").Errf("Please enter the time in 24-hour format (HH:MM)")
func (r *TimeFormatRule) Errf(format string, args ...any) *TimeFormatRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// DateTimeFormatRule validates that a string matches a specified datetime format.
// The format should follow Go's time format specification.
//
// Example:
//
//	rule := DateTimeFormat("2006-01-02 15:04:05").Err("DateTime must be in YYYY-MM-DD HH:MM:SS format")
//	err := rule.Validate("2023-12-31 14:30:00")  // returns nil
//	err = rule.Validate("12/31/2023 2:30 PM")  // returns error
type DateTimeFormatRule struct {
	format string
	e      error
}

// DateTimeFormat creates a new datetime format validation rule.
// The format should follow Go's time format specification.
//
// Example:
//
//	rule := DateTimeFormat("2006-01-02 15:04:05")  // YYYY-MM-DD HH:MM:SS
//	rule := DateTimeFormat("02/01/2006 03:04 PM")  // DD/MM/YYYY HH:MM AM/PM
func DateTimeFormat(format string) *DateTimeFormatRule {
	return &DateTimeFormatRule{
		format: format,
		e:      ErrDateTimeFormat,
	}
}

// Validate checks if the given string matches the specified datetime format.
// Empty strings are considered valid.
//
// Example:
//
//	rule := DateTimeFormat("2006-01-02 15:04:05")
//	err := rule.Validate("2023-12-31 14:30:00")  // returns nil
//	err = rule.Validate("12/31/2023 2:30 PM")  // returns error
//	err = rule.Validate("")  // returns nil (empty string is valid)
func (r *DateTimeFormatRule) Validate(value string) error {
	if value == "" {
		return nil
	}
	_, err := time.Parse(r.format, value)
	if err != nil {
		if r.e != nil {
			return r.e
		}
		return ErrDateTimeFormat
	}
	return nil
}

// Errf sets a custom error message for datetime format validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := DateTimeFormat("2006-01-02 15:04:05").Errf("Please enter the date and time in YYYY-MM-DD HH:MM:SS format")
func (r *DateTimeFormatRule) Errf(format string, args ...any) *DateTimeFormatRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// WeekendRule validates that a time falls on a weekend day (Saturday or Sunday).
//
// Example:
//
//	rule := Weekend().Errf("Event must be on a weekend")
//	err := rule.Validate(time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC))  // returns nil (Saturday)
//	err = rule.Validate(time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC))  // returns nil (Sunday)
//	err = rule.Validate(time.Date(2023, 12, 29, 0, 0, 0, 0, time.UTC))  // returns error (Friday)
type WeekendRule struct {
	e error
}

// Weekend creates a new weekend validation rule.
// The rule ensures that a time value falls on a Saturday or Sunday.
//
// Example:
//
//	rule := Weekend()
func Weekend() *WeekendRule {
	return &WeekendRule{
		e: ErrWeekend,
	}
}

// Validate checks if the given time falls on a weekend day (Saturday or Sunday).
//
// Example:
//
//	rule := Weekend()
//	err := rule.Validate(time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC))  // returns nil (Saturday)
//	err = rule.Validate(time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC))  // returns nil (Sunday)
//	err = rule.Validate(time.Date(2023, 12, 29, 0, 0, 0, 0, time.UTC))  // returns error (Friday)
func (r *WeekendRule) Validate(value time.Time) error {
	weekday := value.Weekday()
	if weekday != time.Saturday && weekday != time.Sunday {
		if r.e != nil {
			return r.e
		}
		return ErrWeekend
	}
	return nil
}

// Errf sets a custom error message for weekend validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Weekend().Errf("This event is only available on weekends")
func (r *WeekendRule) Errf(format string, args ...any) *WeekendRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// WorkdayRule validates that a time falls on a workday (Monday through Friday).
//
// Example:
//
//	rule := Workday().Err("Appointment must be on a workday")
//	err := rule.Validate(time.Date(2023, 12, 29, 0, 0, 0, 0, time.UTC))  // returns nil (Friday)
//	err = rule.Validate(time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC))  // returns error (Saturday)
type WorkdayRule struct {
	e error
}

// Workday creates a new workday validation rule.
// The rule ensures that a time value falls on a Monday through Friday.
//
// Example:
//
//	rule := Workday()
func Workday() *WorkdayRule {
	return &WorkdayRule{
		e: ErrWorkday,
	}
}

// Validate checks if the given time falls on a workday (Monday through Friday).
//
// Example:
//
//	rule := Workday()
//	err := rule.Validate(time.Date(2023, 12, 29, 0, 0, 0, 0, time.UTC))  // returns nil (Friday)
//	err = rule.Validate(time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC))  // returns error (Saturday)
//	err = rule.Validate(time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC))  // returns error (Sunday)
func (r *WorkdayRule) Validate(value time.Time) error {
	weekday := value.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		if r.e != nil {
			return r.e
		}
		return ErrWorkday
	}
	return nil
}

// Errf sets a custom error message for workday validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Workday().Errf("Appointments are only available on weekdays")
func (r *WorkdayRule) Errf(format string, args ...any) *WorkdayRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// HolidayRule validates that a time falls on one of the specified holidays.
//
// Example:
//
//	christmas := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)
//	newYear := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//	rule := Holiday(christmas, newYear).Errf("Date must be a holiday")
//	err := rule.Validate(christmas)  // returns nil
//	err = rule.Validate(newYear)    // returns nil
//	err = rule.Validate(time.Now()) // returns error if not a holiday
type HolidayRule struct {
	holidays []time.Time
	e        error
}

// Holiday creates a new holiday validation rule.
// The rule ensures that a time value falls on one of the specified holidays.
//
// Example:
//
//	christmas := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)
//	newYear := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//	rule := Holiday(christmas, newYear)
func Holiday(holidays ...time.Time) *HolidayRule {
	return &HolidayRule{
		holidays: holidays,
		e:        ErrHoliday,
	}
}

// Validate checks if the given time falls on one of the specified holidays.
// The comparison is done by checking if the year, month, and day match.
//
// Example:
//
//	christmas := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)
//	newYear := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//	rule := Holiday(christmas, newYear)
//	err := rule.Validate(christmas)  // returns nil
//	err = rule.Validate(newYear)    // returns nil
//	err = rule.Validate(time.Date(2023, 12, 26, 0, 0, 0, 0, time.UTC))  // returns error
func (r *HolidayRule) Validate(value time.Time) error {
	valueYear, valueMonth, valueDay := value.Date()
	for _, holiday := range r.holidays {
		holidayYear, holidayMonth, holidayDay := holiday.Date()
		if valueYear == holidayYear && valueMonth == holidayMonth && valueDay == holidayDay {
			return nil
		}
	}
	if r.e != nil {
		return r.e
	}
	return ErrHoliday
}

// Errf sets a custom error message for holiday validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	christmas := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)
//	newYear := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//	rule := Holiday(christmas, newYear).Errf("This service is only available on holidays")
func (r *HolidayRule) Errf(format string, args ...any) *HolidayRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
