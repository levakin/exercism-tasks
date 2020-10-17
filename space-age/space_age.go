// Package space provides functionality for calculating age on different planets.
package space

/*
   Mercury: orbital period 0.2408467 Earth years
   Venus: orbital period 0.61519726 Earth years
   Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
   Mars: orbital period 1.8808158 Earth years
   Jupiter: orbital period 11.862615 Earth years
   Saturn: orbital period 29.447498 Earth years
   Uranus: orbital period 84.016846 Earth years
   Neptune: orbital period 164.79132 Earth years
*/
type Planet string

const (
	SecondsInEarthYear = 31557600

	MercuryOrbitalPeriod = 0.2408467
	VenusOrbitalPeriod   = 0.61519726
	EarthOrbitalPeriod   = 1.0
	MarsOrbitalPeriod    = 1.8808158
	JupiterOrbitalPeriod = 11.862615
	SaturnOrbitalPeriod  = 29.447498
	UranusOrbitalPeriod  = 84.016846
	NeptuneOrbitalPeriod = 164.79132

	// Planets
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Earth   Planet = "Earth"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

// Age calculates age on different planets.
func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / SecondsInEarthYear

	orbitalPeriod := EarthOrbitalPeriod
	switch planet {
	case Mercury:
		orbitalPeriod = MercuryOrbitalPeriod
	case Venus:
		orbitalPeriod = VenusOrbitalPeriod
	case Earth:
		orbitalPeriod = EarthOrbitalPeriod
	case Mars:
		orbitalPeriod = MarsOrbitalPeriod
	case Jupiter:
		orbitalPeriod = JupiterOrbitalPeriod
	case Saturn:
		orbitalPeriod = SaturnOrbitalPeriod
	case Uranus:
		orbitalPeriod = UranusOrbitalPeriod
	case Neptune:
		orbitalPeriod = NeptuneOrbitalPeriod
	}

	return earthYears / orbitalPeriod
}
