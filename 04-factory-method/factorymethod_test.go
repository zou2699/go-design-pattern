/*
@Time : 2019/11/11 15:40
@Author : Tux
@File : factorymethod_test.go
@Description :
*/

package factorymethod

import (

	"testing"
)

func compute(factory OperationFunc, a, b int) (result int) {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperation(t *testing.T) {
	var factory OperationFunc
	factory = OperationAdd{}

	result := compute(factory,100,10)

	if result != 110 {
		t.Fatal("error with factory method pattern")
	}

	factory = OperationSub{}
	result = compute(factory,100,10)

	if result != 90 {
		t.Fatal("error with factory method pattern")
	}

}
