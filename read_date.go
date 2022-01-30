package dsvreader

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var zeroTime time.Time

// Date returns the next date column value from the current row.
//
// date must be in the format YYYY-MM-DD
func (tr *Reader) Date() time.Time {
	if tr.err != nil {
		return zeroTime
	}
	b, err := tr.nextCol()
	if err != nil {
		tr.setColError("cannot read `date`", err)
		return zeroTime
	}
	s := b2s(b)

	y, m, d, err := parseDate(s)
	if err != nil {
		tr.setColError("cannot parse `date`", err)
		return zeroTime
	}
	if y == 0 && m == 0 && d == 0 {
		// special case for ClickHouse
		return zeroTime
	}
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

// DateTime returns the next datetime column value from the current row.
//
// datetime must be in the format YYYY-MM-DD hh:mm:ss.
func (tr *Reader) DateTime() time.Time {
	if tr.err != nil {
		return zeroTime
	}
	b, err := tr.nextCol()
	if err != nil {
		tr.setColError("cannot read `datetime`", err)
		return zeroTime
	}
	s := b2s(b)

	dt, err := parseDateTime(s)
	if err != nil {
		tr.setColError("cannot parse `datetime`", err)
		return zeroTime
	}
	return dt
}

func parseDateTime(s string) (time.Time, error) {
	if len(s) != len("YYYY-MM-DD hh:mm:ss") {
		return zeroTime, errors.New("too short datetime")
	}
	y, m, d, err := parseDate(s[:len("YYYY-MM-DD")])
	if err != nil {
		return zeroTime, err
	}
	s = s[len("YYYY-MM-DD"):]
	if s[0] != ' ' || s[3] != ':' || s[6] != ':' {
		return zeroTime, errors.New("invalid time format. Must be hh:mm:ss")
	}
	hS := s[1:3]
	minS := s[4:6]
	secS := s[7:]
	h, err := strconv.Atoi(hS)
	if err != nil {
		return zeroTime, fmt.Errorf("invalid hour: %w", err)
	}
	min, err := strconv.Atoi(minS)
	if err != nil {
		return zeroTime, fmt.Errorf("invalid minute: %w", err)
	}
	sec, err := strconv.Atoi(secS)
	if err != nil {
		return zeroTime, fmt.Errorf("invalid second: %w", err)
	}
	if y == 0 && m == 0 && d == 0 {
		// Special case for ClickHouse
		return zeroTime, nil
	}
	return time.Date(y, time.Month(m), d, h, min, sec, 0, time.UTC), nil
}

func parseDate(s string) (y, m, d int, err error) {
	if len(s) != len("YYYY-MM-DD") {
		return 0, 0, 0, errors.New("too short date")
	}
	s = s[:len("YYYY-MM-DD")]
	if s[4] != '-' && s[7] != '-' {
		return 0, 0, 0, errors.New("invalid date format. Must be YYYY-MM-DD")
	}
	yS := s[:4]
	mS := s[5:7]
	dS := s[8:]
	y, err = strconv.Atoi(yS)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid year: %w", err)
	}
	m, err = strconv.Atoi(mS)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid month: %w", err)
	}
	d, err = strconv.Atoi(dS)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid day: %w", err)
	}
	return y, m, d, nil
}
