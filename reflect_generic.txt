package main

import (
	"fmt"
	"reflect"
)

// 定义一个泛型结构体 Container
type Container[T any] struct {
	Item T
}

func main() {
	// 创建一个 Container[string] 实例
	container := Container[string]{Item: "initial"}

	// 获取 container 的 reflect.Value
	v := reflect.ValueOf(&container).Elem()

	// 找到 Item 字段
	itemField := v.FieldByName("Item")

	// 确保字段是可设置的
	if itemField.CanSet() {
		// 设置新的值
		itemField.SetString("new value")
	}

	fmt.Println(container.Item) // 应该输出 "new value"
}
