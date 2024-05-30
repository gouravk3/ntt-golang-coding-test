package service

import "github.com/gouravk3/ntt-golang-coding-test/internal/store"

type FuelEstimatorService interface {
	CalculateEstimatedFuel(crewCount int, exoplanet store.Exoplanet) float64
}

type fuelEstimatorService struct {
}

func NewFuelEstimatorService() *fuelEstimatorService {
	return &fuelEstimatorService{}
}
func (fs *fuelEstimatorService) CalculateEstimatedFuel(crewCapacity int, planet store.Exoplanet) float64 {
	// Calculate gravity
	var gravity float64
	if planet.Type == "GasGiant" {
		gravity = 0.5 / (planet.Radius * planet.Radius)
	} else if planet.Type == "Terrestrial" {
		gravity = planet.Mass / (planet.Radius * planet.Radius)
	}

	// Calculate fuel estimation

	fuelEstimation := (float64(planet.Distance) / (gravity * gravity)) * float64(crewCapacity)
	return fuelEstimation

}
