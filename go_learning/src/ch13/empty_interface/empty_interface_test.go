package empty_interface_test

import "testing"

// 断言
func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	println("Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	println("String", s)
	// 	return
	// }
	// println("Unknow Type")

	switch v := p.(type) {
	case int:
		println("Integer", v)
	case string:
		println("String", v)
	default:
		println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("hello")
	DoSomething(10.1)
}
