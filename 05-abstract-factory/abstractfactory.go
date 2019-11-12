/*
@Time : 2019/11/11 16:11
@Author : Tux
@File : abstractfactory
@Description :
Abstract Factory 抽象工厂模式：
          提供一个创建一系列相关或者相互依赖对象的接口，而无需指定他们具体的类
 个人想法：工厂模式和抽象工厂模式：感觉抽象工厂可以叫集团模式，工厂模式下，
         是一个工厂下，对产品的每一个具体生成分配不同的流水线；
         集团模式：在集团下，有不同的工厂，可以生成不同的产品，
         每个工厂生产出来的同一个型号产品具体细节是不一样
*/

// todo
package abstractfactory

import (
	"fmt"
)

// OrderMainDAO 为订单记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory 抽象模式工程接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// RDBMainDAO 为关系型数据库的OrderMainDAO的实现
type RDBMainDAO struct {}

// SaveOrderMain
func (R *RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main save")
}

// RDBDetailDAO  为关系型数据库的OrderDetailDAO的实现
type RDBDetailDAO struct {}

// SaveOrderDetail
func (R *RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save")

}

// RDBDAOFactory 是RDB抽象工厂实现
type RDBDAOFactory struct {}

func (R *RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (R *RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO XML存储
type XMLMainDAO struct {}

func (X *XMLMainDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}

// XMLDetailDAO XML存储
type XMLDetailDAO struct {}

func (X *XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("xml detail save")
}

// XMLDAOFactory 是rdb抽象工厂实现
type XMLDAOFactory struct {}

func (X *XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (X *XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}



