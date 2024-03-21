package core

import "errors"

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (ap *Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}
func (ap *Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}
func (ap *Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}
func (ap *Adapter) Division(a int32, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("the dividend cannot be 0")
	}
	return a + b, nil
}
