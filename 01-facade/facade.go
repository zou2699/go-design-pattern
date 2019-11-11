/*
@Time : 2019/11/11 10:29
@Author : Tux
@File : facade
@Description :
  Facade 外观模式：
        为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，
		这个接口使得这一子系统更加容易使用（投资：基金，股票，房产）
 个人想法：中介者模式、外观模式：每个对象都保存一份中介者对象，
        在和其他对象交互时，通过中介者来完成，外观模式：外观中保存了一堆对象，
		这些对象或者是组成某个子系统的，将其封装在外观对象中，给客户端一种只有一个对象的感觉，
		一个是结构型模式，一个是行为性模式
*/

package facade

import (
	"fmt"
)

// API is facade interface of facade package
type API interface {
	Say() string
}

// AModuleAPI ...
type AModuleAPI interface {
	SayA() string
}

type BModuleAPI interface {
	SayB() string
}

// facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Say() string {
	aRet := a.a.SayA()
	bRet := a.b.SayB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewAPI() API  {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type aModuleImpl struct {
}

func (a aModuleImpl) SayA() string {
	return "A module running"
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type bModuleImpl struct {
}

func (b bModuleImpl) SayB() string {
	return "B module running"
}

// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}
