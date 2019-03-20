package main

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestFirstTry(t *testing.T) {
	t.Log("this is first test")
}

func TestBuildData(t *testing.T) {
	a, b := 1, 1
	fmt.Println(a)
	for i := 0; i < 5; i++ {
		fmt.Println(b)
		a, b = b, a+b
	}
}

func TestIota(t *testing.T) {
	const (
		Monday = iota
		Tuesday
		Wednesday
	)
	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	t.Log(Wednesday, Monday, Tuesday)
	a := 7
	t.Log(Readable&a == Readable, Writable&a == Writable, Executable&a == Executable)
}

func TestTypeSwitch(t *testing.T) {
	type myInt int
	var a int32
	var b int64
	a = 1
	b = 2
	a = int32(b)
	t.Log(a, b)
	var c myInt
	c = myInt(a)
	t.Log(c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Logf("%T *%s*", s, s)
	t.Log(len(s)) //0
}

func TestArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 3, 4}
	//c := [...]int{1,3,3,4,5}
	t.Log(a == b)
	// t.Log(a == c) 长度不同编译报错
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("it is not in 0-3")
		}
	}
}

// 测试 struct json 生成
func TestStructJson(t *testing.T) {
	type IT struct {
		Company  string   `json:"-"`        // 此字段不会输出到屏幕
		Subjects []string `json:"subjects"` // 二次编码
		IsOk     bool     `json:",string"`  // 修改为字符串
		Price    float64  `json:"price"`
	}
	s := IT{"it", []string{"Go", "c++"}, true, 66.66}
	if buf, err := json.MarshalIndent(s, "", " "); err != nil {
		t.Log(err)
	} else {
		t.Log(string(buf))
	}
}

// 测试 map json 生成
func TestMapJson(t *testing.T) {
	m := make(map[string]interface{})
	m["company"] = "cast"
	m["subjects"] = []string{"Go", "c++"}
	m["isOk"] = true
	m["price"] = 66.66

	res, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(string(res))
	}

}

// 测试 json 解析到 结构体
func TestParseStruct(t *testing.T) {
	type IT struct {
		Company  string   `json:"company"`  // 此字段不会输出到屏幕
		Subjects []string `json:"subjects"` // 二次编码
		IsOk     bool     `json:",string"`  // 修改为字符串
		Price    float64  `json:"price"`
	}

	jsonBuf := `
		{
         "company": "cast",
         "isOk": "true",
         "price": 66.66,
         "subjects": [
          "Go",
          "c++"
         ]
        }
	`
	var tem IT
	err := json.Unmarshal([]byte(jsonBuf), &tem) // 第二个参数为引用
	if err != nil {
		t.Log(err)
	} else {
		t.Log(tem)
	}
	// 只解析单个数据
	type IT2 struct {
		Company string `json:"company"`
	}
	var tem2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &tem2)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(tem2)
	}
}

// 测试 json 解析到 map
func TestParseMap(t *testing.T) {

	jsonBuf := `
		{
         "company": "cast",
         "isOk": "true",
         "price": 66.66,
         "subjects": [
          "Go",
          "c++"
         ]
        }
	`
	// 创建map
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(m)
	}
	var str string
	//str = m["company"]
	t.Log(str)
	// 类型断言
	for key, value := range m {
		t.Logf("%v =======> %v", key, value)
		switch data := value.(type) {
		case string:
			str = data
			t.Logf("map[%s] ===> %T ",key,str)
		}
	}
	// 使用 map 但是确定类型要使用类型断言，反推太麻烦了
}
