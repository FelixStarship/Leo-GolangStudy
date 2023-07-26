package main

import "fmt"

// 1.面向对象的设计原则
/*
  1.单一职责
  2.开闭原则
  3.里氏代换原则
  4.依赖倒转原则
  5.接口隔离原则
  6.合成复用原则
  7.迪米特法则
*/

// 1.开闭原则
type AbstractBanker interface {
	DoBusi()
}

type SaveBanker struct {
}

func (s *SaveBanker) DoBusi() {
	fmt.Println("进行存款业务")
}

type TransferBanker struct {
}

func (s *TransferBanker) DoBusi() {
	fmt.Println("进行转账业务")
}

type PayBanker struct {
}

func (p *PayBanker) DoBusi() {
	fmt.Println("进行支付业务")
}

/*
//实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
func BankerBusiness(banker AbstractBanker) {
	//通过接口来向下调用，(多态现象)
	banker.DoBusi()
}
*/

func BankerBussiness(banker ...AbstractBanker) {
	// 多态（调用具体实现层）
	for i := range banker {
		banker[i].DoBusi()
	}
}

func main() {

	BankerBussiness(&SaveBanker{}, &TransferBanker{}, &PayBanker{})
	// 存款
	s := &SaveBanker{}
	s.DoBusi()
	// 转账
	t := &TransferBanker{}
	t.DoBusi()
	// 支付
	p := &PayBanker{}
	p.DoBusi()
}
