package space

// Planet should have a comment documenting it.
type Planet = string

var earthYears float64 = 31557600

// Age should have a comment documenting it.
func Age(seconds float64, planet string) float64 {

	switch planet {
	case "Mercury":
		return seconds / (earthYears * 0.2408467)

	case "Venus":
		return seconds / (earthYears * 0.61519726)

	case "Mars":
		return seconds / (earthYears * 1.8808158)

	case "Jupiter":
		return seconds / (earthYears * 11.862615)

	case "Saturn":
		return seconds / (earthYears * 29.447498)

	case "Uranus":
		return seconds / (earthYears * 84.016846)

	case "Neptune":
		return seconds / (earthYears * 164.79132)

	default:
		return seconds / earthYears
	}

}
