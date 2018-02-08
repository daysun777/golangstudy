package reflection

import (
	"fmt"
	"reflect"
)

type User2 struct {
	//User字段
	Id int
	Name string
	Age int
}

type Manager struct {
	User2
	title string
}

func init3()  {
	m:=Manager{User2: User2{1, "ok", 12}, title:"123"}
	t:=reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(0)) //User2
	//reflect.StructField{Name:"User2", PkgPath:"", Type:(*reflect.rtype)(0x4b7080), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
	fmt.Printf("%#v\n", t.Field(1)) //title
	//reflect.StructField{Name:"title", PkgPath:"gostudy/reflection", Type:(*reflect.rtype)(0x4a8a40), Tag:"", Offset:0x20, Index:[]int{1}, Anonymous:false}

	// 取出User2中当的Name
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,1}))
	//reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4a8b40), Tag:"", Offset:0x8, Index:[]int{1}, Anonymous:false}

}
