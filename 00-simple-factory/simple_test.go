/*
@Time : 2019/11/11 10:21
@Author : Tux
@File : simple_test.go
@Description :
*/

package simple

import (
	"testing"
)

// TestType1 test get hiapi with factory
func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tux")
	if s != "Hi, Tux" {
		t.Fatal("Type1 test fail")
	}
}


// TestType2 test get helloapi with factory
func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tux")
	if s != "Hello, Tux" {
		t.Fatal("Type2 test fail")
	}
}