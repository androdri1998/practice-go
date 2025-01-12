package main

import "fmt"

func buy(job1, job2 bool) (bool, bool, bool) {
	buyLargeTV := job1 && job2
	buySmallTV := job1 != job2
	buyIcecream := job1 || job2

	return buyLargeTV, buySmallTV, buyIcecream
}

func main() {
	tvLarger, tvSmall, iceCream := buy(true, true)
	fmt.Printf("Tv Larger: %t, Tv small: %t, Ice cream: %t, health: %t\n", tvLarger, tvSmall, iceCream, !iceCream)

	tvLarger, tvSmall, iceCream = buy(false, true)
	fmt.Printf("Tv Larger: %t, Tv small: %t, Ice cream: %t, health: %t\n", tvLarger, tvSmall, iceCream, !iceCream)

	tvLarger, tvSmall, iceCream = buy(false, false)
	fmt.Printf("Tv Larger: %t, Tv small: %t, Ice cream: %t, health: %t\n", tvLarger, tvSmall, iceCream, !iceCream)

	tvLarger, tvSmall, iceCream = buy(true, true)
	fmt.Printf("Tv Larger: %t, Tv small: %t, Ice cream: %t, health: %t\n", tvLarger, tvSmall, iceCream, !iceCream)
}
