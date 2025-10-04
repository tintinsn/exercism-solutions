package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	oneYearInSecond := 31557600
	
	switch planet {
	case "Mercury":
		return 	seconds/ float64(oneYearInSecond) / 0.2408467
	case "Venus":
		return 	seconds/ float64(oneYearInSecond) / 0.61519726
	case "Earth":
		return 	seconds/ float64(oneYearInSecond) / 1.0
	case "Mars":
		return 	seconds/ float64(oneYearInSecond) / 1.8808158
	case "Jupiter":
		return 	seconds/ float64(oneYearInSecond) / 11.862615
	case "Saturn":
		return 	seconds/ float64(oneYearInSecond) / 29.447498
	case "Uranus":
		return 	seconds/ float64(oneYearInSecond) / 84.016846
	case "Neptune":
		return 	seconds/ float64(oneYearInSecond) / 164.79132
	} 

	return -1.00
}
