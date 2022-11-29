package main

import (
	"fmt"
	d "wether_request/data"
	w "wether_request/data"
)

func main() {

	var city string
	fmt.Scan(&city)

	x, y, c := d.Data(city)
	fmt.Println(x, y, c)

	s := w.Wetherr(x, y)

	//s = 5 / 9 * (s - 32)

	fmt.Printf("Температура в городе - %s : %0.2fС", c, s)

}
