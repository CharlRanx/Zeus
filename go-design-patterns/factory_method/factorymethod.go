package factorymethod

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂类接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的方法，封装公共的方法
type OperatorBase struct {
	a, b int
}

func (ob *OperatorBase) SetA(a int) {
	ob.a = a
}

func (ob *OperatorBase) SetB(b int) {
	ob.b = b
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

func (po PlusOperator) Result() int {
	return po.a + po.b
}

// MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (mo MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

func (mo MinusOperator) Result() int {
	return mo.a - mo.b
}
