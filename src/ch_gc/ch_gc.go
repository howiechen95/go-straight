package ch_gc

import (
	"fmt"
	"runtime"
)

func PrintStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

// 使用 slice
type data struct {
	i, j int
}

func SliceGC() {
	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}
	runtime.GC()
	_ = structure[0] // 防止垃圾回收器过早地垃圾收集结构体变量，因为未在for循环之外对其进行引用或使用
}

// 使用 pointers 操作 map
func MapWithPointGC() {
	var N = 40000000
	myMap := make(map[int]*int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = &value
	}
	runtime.GC()
	_ = myMap[0]
}

// 不使用 pointers 操作 map
func MapGC() {
	var N = 40000000
	myMap := make(map[int]int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = value
	}
	runtime.GC()
	_ = myMap[0]
}

// 分割 map
func MapSplitGC() {
	var N = 40000000
	split := make([]map[int]int, 200)
	for i := range split {
		split[i] = make(map[int]int)
	}
	for i := 0; i < N; i++ {
		value := int(i)
		split[i%200][value] = value
	}
	runtime.GC()
	_ = split[0][0]
}
