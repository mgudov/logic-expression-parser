package lep

type StartsWithX struct {
	Param *ParamX
	Value Stringify
}

var _ Expression = (*StartsWithX)(nil)
var _ Statement = (*StartsWithX)(nil)

func StartsWith(param *ParamX, value Stringify) *StartsWithX {
	return &StartsWithX{
		Param: param,
		Value: value,
	}
}

func (e StartsWithX) GetParam() *ParamX {
	return e.Param
}

func (e StartsWithX) GetValue() Value {
	return e.Value
}

func (e StartsWithX) Equals(other Expression) bool {
	if expr, ok := other.(*StartsWithX); ok {
		return e.Param.Equals(expr.Param) && e.Value.Equals(expr.Value)
	}
	return false
}

func (e StartsWithX) String() string {
	return e.Param.String() + " starts_with " + e.Value.String()
}

func parseStartsWith(left, right interface{}) (*StartsWithX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	stringify, ok := value.(Stringify)
	if !ok {
		return nil, IncorrectType("parseStartsWith", (Stringify)(nil), right)
	}
	return StartsWith(param, stringify), nil
}

type EndsWithX struct {
	Param *ParamX
	Value Stringify
}

var _ Expression = (*EndsWithX)(nil)
var _ Statement = (*EndsWithX)(nil)

func EndsWith(param *ParamX, value Stringify) *EndsWithX {
	return &EndsWithX{
		Param: param,
		Value: value,
	}
}

func (e EndsWithX) GetParam() *ParamX {
	return e.Param
}

func (e EndsWithX) GetValue() Value {
	return e.Value
}

func (e EndsWithX) Equals(other Expression) bool {
	if expr, ok := other.(*EndsWithX); ok {
		return e.Param.Equals(expr.Param) && e.Value.Equals(expr.Value)
	}
	return false
}

func (e EndsWithX) String() string {
	return e.Param.String() + " ends_with " + e.Value.String()
}

func parseEndsWith(left, right interface{}) (*EndsWithX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	stringify, ok := value.(Stringify)
	if !ok {
		return nil, IncorrectType("parseEndsWith", (Stringify)(nil), right)
	}
	return EndsWith(param, stringify), nil
}
