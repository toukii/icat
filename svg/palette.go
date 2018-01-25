package svg

import (
	"fmt"
	"math"
)

type C struct {
	R, G, B uint8
	Θ       float64
}

func GetC2(du int) *C {
	return GetC(float64(du) * math.Pi / 180)
}

func GetC(θ float64) *C {
	c := new(C)
	c.SetC(θ)
	return c
}

func (c *C) SetC(θ float64) {
	c.Θ = θ
	c.Compute()
}

func (c *C) Compute() {
	c.R = uint8(255 * math.Sin(c.Θ/2))
	c.G = uint8(255 * math.Sin(c.Θ/2+math.Pi/3))
	c.B = uint8(255 * math.Sin(c.Θ/2+math.Pi*2/3))
}

func (c *C) String() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}
