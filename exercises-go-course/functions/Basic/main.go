package main

import "fmt"

func f1() {
	fmt.Println("Calling function f1")
}

func f2(p1 string, p2 int) {
	fmt.Println(p1, p2)
}

func f3() string {
	return "Something"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("%s %s", p1, p2)
}

func f5() (string, string) {
	return "New", "Something"
}

func main() {
	f1()
	f2("Something", 100)
	fmt.Println(f3())
	fmt.Println(f4("Other", "Something"))
	s1, s2 := f5()
	fmt.Println(s1, s2)
}
