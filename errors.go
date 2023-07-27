package lep

import "fmt"

type ErrIncorrectType struct {
	FuncName string
	Expected interface{}
	Received interface{}
}

func IncorrectType(funcName string, expected, received interface{}) error {
	return ErrIncorrectType{
		FuncName: funcName,
		Expected: expected,
		Received: received,
	}
}

func (e ErrIncorrectType) Error() string {
	return fmt.Sprintf("%s: incorrect type; expected: %T; received: %T", e.FuncName, e.Expected, e.Received)
}

type ErrIncorrectValue struct {
	FuncName string
	Expected interface{}
	Received interface{}
}

func IncorrectValue(funcName string, expected, received interface{}) error {
	return ErrIncorrectValue{
		FuncName: funcName,
		Expected: expected,
		Received: received,
	}
}

func (e ErrIncorrectValue) Error() string {
	return fmt.Sprintf("%s: incorrect value; expected: %T; received: %T", e.FuncName, e.Expected, e.Received)
}
