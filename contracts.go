package lep

type Expression interface {
	Equals(Expression) bool
	String() string
}

type Value interface {
	Expression
	Value() interface{}
}
