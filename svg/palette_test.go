package svg

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	c := GetC2(0)
	fmt.Println(0, c)

	c = GetC2(90)
	fmt.Println(90, c)

	c = GetC2(120)
	fmt.Println(120, c)

	c = GetC2(180)
	fmt.Println(180, c)

	c = GetC2(270)
	fmt.Println(270, c)

}
