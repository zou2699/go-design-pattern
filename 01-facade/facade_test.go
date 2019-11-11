/*
@Time : 2019/11/11 14:01
@Author : Tux
@File : facade_test.go
@Description :
*/

package facade

import (
	"testing"
)

var expect = "A module running\nB module running"

func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Say()
	if ret != expect {
		t.Fatalf("except %s, return %s", expect, ret)
	}
}
