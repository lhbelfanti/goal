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
	luxury bool
}

func main () {
	t := truck {
		vehicle: vehicle{
			doors: 2,
			color: "yellow",
		},
		fourWheel: true,
	}

	s := sedan {
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Printf("Truck-> doors: %d, color: %s, fourWheel: %v\n", t.doors, t.color, t.fourWheel)
	fmt.Printf("Sedan-> doors: %d, color: %s, luxury: %v", s.doors, s.color, s.luxury)
}
