/*
@Time : 2019/11/11 14:44
@Author : Tux
@File : adapter_test.go
@Description :
*/

package adapter

import (
	"testing"
)

// var expect = "adaptee method"

func TestNewAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	expect := adaptee.SpecificRequest()
	s := target.Request()
	if s != adaptee.SpecificRequest() {
		t.Fatalf("expect %s, actual %s", expect, s)
	}
}
