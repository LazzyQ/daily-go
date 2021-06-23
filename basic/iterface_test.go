package basic

import "testing"

type TestInterface interface {
	Test()
}

type Implement1 struct{}

func (i Implement1) Test() {

}

type Implement2 struct{}

func (i *Implement2) Test() {

}

func TestInterfaceImplement(t *testing.T) {
	var i1 interface{} = Implement1{}
	var ii1 interface{} = &Implement1{}
	var i2 interface{} = Implement2{}
	var ii2 interface{} = &Implement2{}

	_, ok1 := i1.(TestInterface)
	_, ok11 := ii1.(TestInterface)
	_, ok2 := i2.(TestInterface)
	_, ok22 := ii2.(TestInterface)

	t.Logf("%v,  %v, %v, %v", ok1, ok11, ok2, ok22)
}
