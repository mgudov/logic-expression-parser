package lep

type Equals statement

var _ Expression = (*Equals)(nil)

func NewEquals(param *Param, value Value) *Equals {
	return &Equals{
		Param: param,
		Value: value,
	}
}

func (e Equals) Equals(other Expression) bool {
	if expr, ok := other.(*Equals); ok {
		return e.Param.Equals(expr.Param) && e.Value.Equals(expr.Value)
	}
	return false
}

func (e Equals) String() string {
	return e.Param.String() + "=" + e.Value.String()
}

func parseEquals(left, right interface{}) (*Equals, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return NewEquals(param, value), nil
}

type NotEquals statement

var _ Expression = (*NotEquals)(nil)

func NewNotEquals(param *Param, value Value) *NotEquals {
	return &NotEquals{
		Param: param,
		Value: value,
	}
}

func (e NotEquals) Equals(other Expression) bool {
	if expr, ok := other.(*NotEquals); ok {
		return e.Param.Equals(expr.Param) && e.Value.Equals(expr.Value)
	}
	return false
}

func (e NotEquals) String() string {
	return e.Param.String() + "!=" + e.Value.String()
}

func parseNotEquals(left, right interface{}) (*NotEquals, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return NewNotEquals(param, value), nil
}
