package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxary bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: true,
	}

	c := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxary: true,
	}

	fmt.Printf("%v\t%v\n", t, t.fourWheel)
	fmt.Printf("%v\t%v\n", c, c.luxary)
}
