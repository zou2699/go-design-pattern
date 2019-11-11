/*
@Time : 2019/11/11 15:02
@Author : Tux
@File : singleton_test.go
@Description :
*/

package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 !=instance2 {
		t.Fatal("instance is not equal")
	}
}
