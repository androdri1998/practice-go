package main

import "fmt"

func typeValue(value interface{}) string {
	switch value.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(typeValue(2.3))
	fmt.Println(typeValue(2))
	fmt.Println(typeValue("anything"))
	fmt.Println(typeValue(func() {}))
	fmt.Println(typeValue(true))
}
