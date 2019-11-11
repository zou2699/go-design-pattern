/*
@Time : 2019/11/11 14:58
@Author : Tux
@File : singleton
@Description :
Singleton 单件：
        保证一个类仅有一个实例，并提供一个访问它的全局访问点
 个人想法：用Go实现时，巧妙使用包级别的变量声明规则：小写字母的包级别变量是不对外开放的，
         创建实例时，用同步库sync.Once来保证全局只有一个对象实例
*/

package singleton

import (
	"fmt"
	"sync"
)

// Singleton 是单例模式类
type Singleton struct {
	data int
}

// 定义一个包级别的private实例变量
var singleton *Singleton

// 同步Once,保证每次调用时，只有第一次生效
var once sync.Once

// GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{26}
	})
	fmt.Println("实例对象的信息和地址", singleton, &singleton.data)

	return singleton
}
