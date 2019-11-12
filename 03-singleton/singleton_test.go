/*
@Time : 2019/11/11 15:02
@Author : Tux
@File : singleton_test.go
@Description :
*/

package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	instance1 := GetInstance(1)
	instance2 := GetInstance(2)

	fmt.Println("实例对象的信息和地址", &instance1.data, &instance2.data)
	if instance1 !=instance2 {
		t.Fatal("instance is not equal")
	}
}
