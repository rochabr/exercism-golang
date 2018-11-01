package raindrops

import (
	"fmt"
	"math"
)

//Convert shpuld be
func Convert(number float64) string {
	result := ""
	if math.Mod(number, 3) == 0 {
		result += "Pling"
	}

	if math.Mod(number, 5) == 0 {
		result += "Plang"
	}

	if math.Mod(number, 7) == 0 {
		result += "Plong"
	}

	if result != "" {
		return result
	}

	return fmt.Sprintf("%.f", number)

}
