package type_test

import "testing"

type MyInt int64

// go 无法支持隐式类型转换,别名隐式类型转换也不支持
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

// go语言不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	// 获取指针
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

// string 是值类型，其默认的初始化值为空字符串串，⽽而不不是 nil
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") //初始化零值是“”
	t.Log(len(s))

}
