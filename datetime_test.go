package lep

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseDateTime(t *testing.T) {
	type testParseDateTime struct {
		string Expression
		result time.Time
		err    error
	}
	var tests = []testParseDateTime{
		{
			string: NewString("2020-10-20"),
			result: time.Date(2020, 10, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			string: NewString("2020-10-20 10:20:30"),
			result: time.Date(2020, 10, 20, 10, 20, 30, 0, time.UTC),
		},
		{
			string: NewString("20.10.2020"),
			err:    errors.New(`parsing time "20.10.2020": month out of range`),
		},
		{
			string: NewInteger(1),
			err:    NewIncorrectType("parseDateTime", (*String)(nil), (*Integer)(nil)),
		},
	}

	for _, tt := range tests {
		d, err := parseDateTime(tt.string)
		if tt.err == nil && assert.NoError(t, err) && d != nil {
			assert.IsType(t, (*DateTime)(nil), d)
			assert.Equal(t, tt.result, d.Val)
			assert.Equal(t, tt.result, d.Value())
			assert.Equal(t, "dt:"+tt.string.String(), d.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestDateTime_Equals(t *testing.T) {
	var (
		date1 = time.Date(2020, 10, 20, 10, 20, 30, 0, time.UTC)
		date2 = time.Date(1990, 11, 22, 11, 22, 33, 0, time.UTC)
	)

	type testDateTimeEquals struct {
		d1     Expression
		d2     Expression
		result bool
	}
	var tests = []testDateTimeEquals{
		{
			d1:     NewDateTime(date1, "2006-01-02"),
			d2:     NewDateTime(date1, "2006-01-02"),
			result: true,
		},
		{
			d1:     NewDateTime(date1, "2006-01-02"),
			d2:     NewDateTime(date1, "2006-01-02 15:04:05"),
			result: true,
		},
		{
			d1:     NewDateTime(date1, "2006-01-02 15:04:05"),
			d2:     NewDateTime(date1, "2006-01-02 15:04:05"),
			result: true,
		},
		{
			d1:     NewDateTime(date1, "2006-01-02"),
			d2:     NewDateTime(date2, "2006-01-02"),
			result: false,
		},
		{
			d1:     NewDateTime(date1, "2006-01-02"),
			d2:     NewDateTime(date2, "2006-01-02 15:04:05"),
			result: false,
		},
		{
			d1:     NewDateTime(date1, "2006-01-02 15:04:05"),
			d2:     NewDateTime(date2, "2006-01-02 15:04:05"),
			result: false,
		},
		{
			d1:     NewDateTime(date1, "2006-01-02"),
			d2:     NewString(date1.String()),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.d1.Equals(tt.d2))
		assert.Equal(t, tt.result, tt.d2.Equals(tt.d1))
	}
}
