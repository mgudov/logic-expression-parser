package lep

import (
	"github.com/araddon/dateparse"
	"time"
)

type DateTime struct {
	Val    time.Time
	Format string
}

var _ Value = (*DateTime)(nil)

func NewDateTime(dt time.Time, format string) *DateTime {
	return &DateTime{
		Val:    dt,
		Format: format,
	}
}

func (v DateTime) Equals(other Expression) bool {
	if expr, ok := other.(*DateTime); ok {
		return v.Val == expr.Val
	}
	return false
}

func (v DateTime) String() string {
	return `dt:"` + v.Val.Format(v.Format) + `"`
}

func (v DateTime) Value() interface{} {
	return v.Val
}

func parseDateTime(val interface{}) (*DateTime, error) {
	strVal, ok := val.(*String)
	if !ok {
		return nil, NewIncorrectType("parseDateTime", (*String)(nil), val)
	}
	format, err := dateparse.ParseFormat(strVal.Val)
	if err != nil {
		return nil, err
	}

	dt, _ := time.Parse(format, strVal.Val)
	return NewDateTime(dt, format), nil
}
