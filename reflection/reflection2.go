package reflection

import (
	"fmt"
	"reflect"
)

type user struct {
	//User字段
	Id int
	Name string
	Age int
}

//User方法
func (u user) Hello()  {
	fmt.Println("hello world")
}

// 反射的基本操作
func info(o interface{})  {
	//取类型
	t:=reflect.TypeOf(o)
	fmt.Println("Type: ", t.Name())

	if k:=t.Kind();k!=reflect.Struct {
		fmt.Println("类型错误")
		return
	}

	//取值
	v:=reflect.ValueOf(o)
	fmt.Println("Fields:")

	// 打印字段
	for i:=0;i<t.NumField();i++{
		f:=t.Field(i)				//取到字段
		val:=v.Field(i).Interface() //取到字段对应的值
		fmt.Printf("%6s: %v = %v \n", f.Name, f.Type, val)
	}

	/**
	Type:  User
	Fields:
		Id: int = 1
	  Name: string = ok
	   Age: int = 12
	 */

	 // 打印方法
	 for i:=0;i<t.NumMethod();i++{
	 	m:=t.Method(i)
	 	fmt.Printf("%6s: %v\n", m.Name, m.Type) // Hello: func(reflection.User)
	 }

}

func init2()  {
	u:=user{1, "ok", 12}
	info(u)
	info(100) //类型错误
}
