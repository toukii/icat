package svg

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	c := GetC2(0)
	fmt.Println(0, c.String())

	c = GetC2(120)
	fmt.Println(120, c.String())

	c = GetC2(240)
	fmt.Println(240, c.String())

	c = GetC2(90)
	fmt.Println(90, c.String())

	c = GetC2(180)
	fmt.Println(180, c.String())

}
