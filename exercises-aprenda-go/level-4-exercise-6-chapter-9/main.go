package main

import "fmt"

func main() {
	list := make([]string, 26, 26)
	list = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(list))
	fmt.Println(cap(list))

	for index := 0; index < len(list); index++ {
		fmt.Println(list[index])
	}
}
