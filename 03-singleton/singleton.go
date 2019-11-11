/*
@Time : 2019/11/11 14:58
@Author : Tux
@File : singleton
@Description :
*/

package singleton

import (
	"sync"
)

// Singleton 是单例模式类
type Singleton struct {
}

var singleton *Singleton
var once sync.Once

// GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
