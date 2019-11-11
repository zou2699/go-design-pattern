/*
@Time : 2019/11/11 15:11
@Author : Tux
@File : factorymethod
@Description :
Factory Method 工厂方法模式：
        定义一个用于创建对象的接口，让子类决定实例化哪一个类。
        工厂方法使一个类的实例化延迟到其子类

 个人想法：简单工厂和工厂模式：简单工厂定义的是静态函数，
        一个函数处理所有的产品创建，工厂模式将创建对象过程抽象为一个类组，
		有抽象类，有对应产品的创建类，创建的过程有创建类来完成，
		工厂模式主要使用的是依赖反转原则
		（高层模块不依赖底层模块，统一依赖抽象层，抽象层不依赖细节层，细节层依赖抽象层），
		解决简单工厂的缺少开放-封闭原则
*/

package factorymethod

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type Operation struct {
	a int
	b int
}

// SetA 设置 A
func (o *Operation) SetA(a int) {
	o.a = a
}

// SetB 设置 B
func (o *Operation) SetB(b int) {
	o.b = b
}

type OperationAdd struct {
	Operation
}

func (o OperationAdd) Result() int {
	return o.a + o.b
}

type OperationSub struct {
	Operation
}

func (o OperationSub) Result() int {
	return o.a - o.b
}

// OperationFunc 是工厂接口
type OperationFunc interface {
	Create() Operator
}

func (OperationAdd) Create() Operator  {
	return &OperationAdd{}
}

func (OperationSub) Create() Operator  {
	return &OperationSub{}
}