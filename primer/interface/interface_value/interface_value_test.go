package interface_value

import (
	"fmt"
	"math"
)

// 接口也是值。它们可以像其它值一样传递。

// 接口值可以用作函数的参数或返回值。

// 在内部，接口值可以看做包含值和具体类型的元组：

// (value, type)

// 接口值保存了一个具体底层类型的具体值。

// 接口值调用方法时会执行其底层类型的同名方法。

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func Example_interface_value() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
	// Output:
	// (&{Hello}, *interface_value.T)
	// Hello
	// (3.141592653589793, interface_value.F)
	// 3.141592653589793
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
