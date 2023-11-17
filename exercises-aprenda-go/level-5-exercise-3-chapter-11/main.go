package main

import "fmt"

type veicle struct {
	doors int
	color string
}

type truck struct {
	veicle
	traction bool
}

type sedan struct {
	veicle
	luxModel bool
}

func main() {
	truck1 := truck{veicle{2, "red"}, true}
	sedan1 := sedan{veicle{4, "blue"}, false}

	fmt.Println(truck1)
	fmt.Println(truck1.color)
	fmt.Println(truck1.veicle.color)

	fmt.Println(sedan1)
	fmt.Println(sedan1.color)
	fmt.Println(sedan1.veicle.color)
}
