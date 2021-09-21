package lep

import "regexp"

type RegexpX struct {
	Regexp *regexp.Regexp
}

var _ Expression = (*RegexpX)(nil)

func Regexp(regexp *regexp.Regexp) *RegexpX {
	return &RegexpX{Regexp: regexp}
}

func (e RegexpX) Equals(other Expression) bool {
	if expr, ok := other.(*RegexpX); ok {
		return e.Regexp.String() == expr.Regexp.String()
	}
	return false
}

func (e RegexpX) String() string {
	return e.Regexp.String()
}

func parseRegexp(b []byte) (*RegexpX, error) {
	re, err := regexp.Compile(string(b))
	if err != nil {
		return nil, err
	}
	return Regexp(re), nil
}

type MatchRegexpX struct {
	Param  *ParamX
	Regexp *RegexpX
}

var _ Expression = (*MatchRegexpX)(nil)

func MatchRegexp(param *ParamX, regexp *RegexpX) *MatchRegexpX {
	return &MatchRegexpX{
		Param:  param,
		Regexp: regexp,
	}
}

func (e MatchRegexpX) Equals(other Expression) bool {
	if expr, ok := other.(*MatchRegexpX); ok {
		return e.Param.Equals(expr.Param) && e.Regexp.Equals(expr.Regexp)
	}
	return false
}

func (e MatchRegexpX) String() string {
	return e.Param.String() + " =~ " + e.Regexp.String()
}

func parseMatchRegexp(left, right interface{}) (*MatchRegexpX, error) {
	param, ok := left.(*ParamX)
	if !ok {
		return nil, IncorrectType("parseMatchRegexp", (*ParamX)(nil), left)
	}
	re, ok := right.(*RegexpX)
	if !ok {
		return nil, IncorrectType("parseMatchRegexp", (*RegexpX)(nil), right)
	}
	return MatchRegexp(param, re), nil
}

type NotMatchRegexpX struct {
	Param  *ParamX
	Regexp *RegexpX
}

var _ Expression = (*NotMatchRegexpX)(nil)

func NotMatchRegexp(param *ParamX, regexp *RegexpX) *NotMatchRegexpX {
	return &NotMatchRegexpX{
		Param:  param,
		Regexp: regexp,
	}
}

func (e NotMatchRegexpX) Equals(other Expression) bool {
	if expr, ok := other.(*NotMatchRegexpX); ok {
		return e.Param.Equals(expr.Param) && e.Regexp.Equals(expr.Regexp)
	}
	return false
}

func (e NotMatchRegexpX) String() string {
	return e.Param.String() + " =~ " + e.Regexp.String()
}

func parseNotMatchRegexp(left, right interface{}) (*NotMatchRegexpX, error) {
	param, ok := left.(*ParamX)
	if !ok {
		return nil, IncorrectType("parseNotMatchRegexp", (*ParamX)(nil), left)
	}
	re, ok := right.(*RegexpX)
	if !ok {
		return nil, IncorrectType("parseNotMatchRegexp", (*RegexpX)(nil), right)
	}
	return NotMatchRegexp(param, re), nil
}
