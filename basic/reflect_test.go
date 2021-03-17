package basic

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Age  int
	Name string
}

func (u User) GetAge() int {
	return u.Age
}

func (u *User) GetName() string {
	return u.Name
}

func (u User) Print() {
	println("====")
}

func (u *User) PrintPointer() {
	println("****")
}

func TestNil(t *testing.T) {
	var u *User
	u.PrintPointer()
}

func TestReflectField(t *testing.T) {
	user := User{1, "小强"}

	ut := reflect.TypeOf(user)
	uv := reflect.ValueOf(user)
	t.Log(ut, uv)

	for i := 0; i < ut.NumField(); i++ {
		f := ut.Field(i)
		t.Logf("field name: %s", f.Name)
	}
}

func TestReflectMethod(t *testing.T) {
	user := User{1, "小强"}

	ut := reflect.TypeOf(user)
	uv := reflect.ValueOf(user)
	t.Log(ut, uv)

	for i := 0; i < ut.NumMethod(); i++ {
		m := ut.Method(i)
		t.Logf("methd name: %s", m.Name)
	}
}

func TestReflectMethod2(t *testing.T) {
	user := User{1, "小强"}

	ut := reflect.TypeOf(&user)
	uv := reflect.ValueOf(&user)
	t.Log(ut, uv)

	for i := 0; i < ut.NumMethod(); i++ {
		m := ut.Method(i)
		t.Logf("methd name: %s", m.Name)
	}
}

func TestReflect(t *testing.T) {
	user := User{1, "小强"}

	ut := reflect.TypeOf(user)
	uv := reflect.ValueOf(user)
	t.Log(ut, uv)
	t.Log(uv.Type())

	userBack := uv.Interface().(User)
	t.Log(userBack)
}

type MyInt int

func TestReflectModify(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v.CanSet()) // false
	fmt.Println(v.Kind())   // float64

	v = reflect.ValueOf(&x)
	fmt.Println(v.CanSet()) // false

	fmt.Println(v.Elem().CanSet()) // true

	fmt.Println(v.Kind()) // ptr

	var xx MyInt = 7
	xxi := reflect.ValueOf(xx)
	fmt.Println(xxi.Kind()) // int
}

type Interface1 interface {
	String1()
}

type Interface2 interface {
	String2()
}

type Implements struct {
}

func (i Implements) String1() {
	fmt.Println("String1")
}

func (i Implements) String2() {
	fmt.Println("String2")
}

func TestTypeConvert(t *testing.T) {
	i := Implements{}
	t.Log(reflect.TypeOf(i))

	j := Interface1(i)
	t.Log(reflect.TypeOf(j))

	_, ok := j.(Implements)
	t.Log(ok)

	var a interface{} = 1
	b := a.(int)
	t.Log(b)
}
