package lep

func ParseExpression(data string, opts ...Option) (Expression, error) {
	result, err := Parse("data", []byte(data), opts...)
	if err != nil {
		return nil, err
	}
	return result.(Expression), nil
}
