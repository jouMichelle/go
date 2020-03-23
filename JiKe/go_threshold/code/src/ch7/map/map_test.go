package my_map

import "testing"

// map初始化
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	// 打印map的长度
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	// 为map中新增一个值
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

// map中面对不存在的key map会分配一个默认值，如何进行区分
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	// map取值会返回两个值，第一个为值，第二个为判断key是否存在;true存在、false不存在
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

// map进行遍历
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	// 第一个值代表key，第二个值代表 value
	for k, v := range m1 {
		t.Log(k, v)
	}
}
