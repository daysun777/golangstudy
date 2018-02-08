package structs

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

/**
结构体类型 Person
属性大写外部可以访问，小写内部可以访问
*/
type Person struct {
	Name string
	Age  int
}

// Person类的方法 大写对外可以调用
func (person Person) LogPerson() {
	fmt.Printf("名字：%s，年龄：%d \n", person.Name, person.Age)
}

type User struct {
	Name string
	Age  int32
	mess string
}

type Student struct {
	User  //继承 U User U别名
	Id    int64
	Grade int8
}

//伪造构造函数
func NewUser(name string, age int32, mess string) *User {
	return &User{name, age, mess}
}

func (this *User) init(name string, age int32, mess string) {
	this.Name = name
	this.Age = age
	this.mess = mess
}

func (this User) getName() string {
	return this.Name
}

func testuser() {
	//type 1
	var user *User = &User{
		Name: "Nick",
		Age:  20,
		mess: "lover",
	}

	fmt.Println(user)                    //&{Nick 20 lover}
	fmt.Println(user.mess, (*user).mess) //lover lover

	//type 2
	var user1 User
	user1.Name = "dawn"
	user1.Age = 18
	fmt.Println(user1.Name, user1.Age) //dawn 18

	//type 3
	var user2 *User = new(User)
	user2.Name = "xm"
	user2.Age = 18
	fmt.Println(user2)                     //&{xm 18 }
	fmt.Println(user2.Name, (*user2).Name) //xm xm

	user3 := NewUser("xm", 18, "lover")
	fmt.Println(user3, user3.mess, user3.Name, user3.Age) //&{xm 18 lover} lover xm 18

	fmt.Printf("Nmae:%p\n", &user3.Name) //Nmae:0xc042066120
	fmt.Printf("Age :%p\n", &user3.Age)  //Age :0xc042066130
	fmt.Printf("mess:%p\n", &user3.mess) //mess:0xc042066138
}

func testuser2() {
	var user User
	user.init("nick", 18, "lover")
	name := user.getName()
	fmt.Println(name) //nick
}

func teststudent() {
	var stu Student
	stu.User.Name = "小明"
	stu.User.Age = 10
	stu.Id = 1000
	stu.Grade = 5

	fmt.Println(stu)
	fmt.Println(stu.User.Name)
}

func init() {
	//testuser()
	//testuser2()
	teststudent()
}
