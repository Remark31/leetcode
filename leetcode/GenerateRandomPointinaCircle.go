package leetcode

import "math/rand"

type Solution struct {
	radius   float64
	x_center float64
	y_center float64
	x_min    float64
	y_min    float64
	ranges   float64
}

func ConstructorRandom(radius float64, x_center float64, y_center float64) Solution {
	s := Solution{}
	s.radius = radius
	s.x_center = x_center
	s.y_center = y_center
	s.x_min = x_center - radius
	s.y_min = y_center - radius
	s.ranges = radius * radius
	return s
}

func (this *Solution) RandPoint() []float64 {
	for {
		n1 := rand.Float64()
		n2 := rand.Float64()
		x, y := 2*n1*this.radius+this.x_min, 2*n2*this.radius+this.y_min
		if n1*this.radius*n1*this.radius+n2*this.radius*n2*this.radius <= this.ranges {
			return []float64{x, y}
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
