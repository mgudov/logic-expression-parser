package lep

import (
	"github.com/araddon/dateparse"
	"time"
)

type DateTimeX struct {
	Val    time.Time
	Format string
}

var _ Value = (*DateTimeX)(nil)

func DateTime(dt time.Time, format string) *DateTimeX {
	return &DateTimeX{
		Val:    dt,
		Format: format,
	}
}

func (v DateTimeX) Equals(other Expression) bool {
	if expr, ok := other.(*DateTimeX); ok {
		return v.Val == expr.Val
	}
	return false
}

func (v DateTimeX) String() string {
	return `dt:"` + v.Val.Format(v.Format) + `"`
}

func (v DateTimeX) Value() interface{} {
	return v.Val
}

func parseDateTime(val interface{}) (*DateTimeX, error) {
	strVal, ok := val.(*StringX)
	if !ok {
		return nil, IncorrectType("parseDateTime", (*StringX)(nil), val)
	}
	format, err := dateparse.ParseFormat(strVal.Val)
	if err != nil {
		return nil, err
	}

	dt, _ := time.Parse(format, strVal.Val)
	return DateTime(dt, format), nil
}
