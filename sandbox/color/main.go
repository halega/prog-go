package main

import (
	"fmt"

	"github.com/fatih/color"
)

type coloredGrade int

func (g coloredGrade) String() string {
	switch g {
	case 1:
		return color.RedString("%d", g)
	case 2:
		return color.HiRedString("%d", g)
	case 3:
		return color.HiBlueString("%d", g)
	case 4:
		return color.YellowString("%d", g)
	case 5:
		return color.HiGreenString("%d", g)
	default:
		return fmt.Sprintf("%d", g)
	}
}

func main() {
	fmt.Println(coloredGrade(1))
}
