package reflection

import (
	"reflect"
	"fmt"
)

/**
- 反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
- 反身会将匿名字段作为独立字段（匿名字段本质）
- 想要利用反射修改对象状态，前提是interface.data是settable,即pointer-interface
- 通过反射可以“动态”调用方法
 */

type NotknownType struct {
	S1 string
	S2 string
	S3 string
}

func (n NotknownType) string() string  {
	return n.S1 + " & " + n.S2 + " & " + n.S3
}

var secret interface{} = NotknownType{"Go", "C", "Python"}

func init1()  {
	value := reflect.ValueOf(secret)
	fmt.Println(value) //{Go C Python}

	typ := reflect.TypeOf(secret)
	fmt.Println(typ) //goreflect.NotknownType

	knd := value.Kind()
	fmt.Println(knd)

	for i:=0;i<value.NumField();i++{
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		/**
		Field 0: Go
		Field 1: C
		Field 2: Python
		 */
	}

}