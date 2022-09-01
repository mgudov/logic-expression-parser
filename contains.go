package lep

type HasX struct {
	statement
}

var _ Expression = (*HasX)(nil)
var _ Statement = (*HasX)(nil)

func Has(param *ParamX, value Value) *HasX {
	return &HasX{
		statement: statement{
			Param: param,
			Value: value,
		},
	}
}

func (e HasX) Equals(other Expression) bool {
	if expr, ok := other.(*HasX); ok {
		return e.Param.Equals(expr.Param) && e.Value.Equals(expr.Value)
	}
	return false
}

func (e HasX) String() string {
	return e.Param.String() + " has " + e.Value.String()
}

func parseHas(left, right interface{}) (*HasX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return Has(param, value), nil
}

type NotHasX struct {
	statement
}

var _ Expression = (*NotHasX)(nil)
var _ Statement = (*NotHasX)(nil)

func NotHas(param *ParamX, value Value) *NotHasX {
	return &NotHasX{
		statement: statement{
			Param: param,
			Value: value,
		},
	}
}

func (e NotHasX) Equals(other Expression) bool {
	if expr, ok := other.(*NotHasX); ok {
		return e.Param.Equals(expr.Param) && e.Value.Equals(expr.Value)
	}
	return false
}

func (e NotHasX) String() string {
	return e.Param.String() + " not_has " + e.Value.String()
}

func parseNotHas(left, right interface{}) (*NotHasX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return NotHas(param, value), nil
}

type HasAnyX struct {
	Param *ParamX
	Slice *SliceX
}

var _ Expression = (*HasAnyX)(nil)
var _ Statement = (*HasAnyX)(nil)

func HasAny(param *ParamX, slice *SliceX) *HasAnyX {
	return &HasAnyX{
		Param: param,
		Slice: slice,
	}
}

func (e HasAnyX) Equals(other Expression) bool {
	if expr, ok := other.(*HasAnyX); ok {
		return e.Param.Equals(expr.Param) && e.Slice.Equals(expr.Slice)
	}
	return false
}

func (e HasAnyX) String() string {
	return e.Param.String() + " has_any " + e.Slice.String()
}

func (e HasAnyX) GetParam() *ParamX {
	return e.Param
}

func (e HasAnyX) GetValue() Value {
	return e.Slice
}

func parseHasAny(left, right interface{}) (*HasAnyX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	slice, ok := value.(*SliceX)
	if !ok {
		return nil, IncorrectType("parseHasAny", (*SliceX)(nil), value)
	}
	return HasAny(param, slice), nil
}

type HasAllX struct {
	Param *ParamX
	Slice *SliceX
}

var _ Expression = (*HasAllX)(nil)
var _ Statement = (*HasAllX)(nil)

func HasAll(param *ParamX, slice *SliceX) *HasAllX {
	return &HasAllX{
		Param: param,
		Slice: slice,
	}
}

func (e HasAllX) Equals(other Expression) bool {
	if expr, ok := other.(*HasAllX); ok {
		return e.Param.Equals(expr.Param) && e.Slice.Equals(expr.Slice)
	}
	return false
}

func (e HasAllX) String() string {
	return e.Param.String() + " has_all " + e.Slice.String()
}

func (e HasAllX) GetParam() *ParamX {
	return e.Param
}

func (e HasAllX) GetValue() Value {
	return e.Slice
}

func parseHasAll(left, right interface{}) (*HasAllX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	slice, ok := value.(*SliceX)
	if !ok {
		return nil, IncorrectType("parseHasAll", (*SliceX)(nil), value)
	}
	return HasAll(param, slice), nil
}
