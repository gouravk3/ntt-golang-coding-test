package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gouravk3/ntt-golang-coding-test/config"
	"github.com/gouravk3/ntt-golang-coding-test/internal/service"
	"github.com/gouravk3/ntt-golang-coding-test/internal/store"
)

type fuelEstimatorHandlers struct {
	fuelEstimatorService service.FuelEstimatorService
	config               config.AppConfig
}

func NewfuelEstimatorHandler(config config.AppConfig) *fuelEstimatorHandlers {
	return &fuelEstimatorHandlers{
		config:               config,
		fuelEstimatorService: service.NewFuelEstimatorService(),
	}
}

func (h *fuelEstimatorHandlers) FuelEstimation(c *gin.Context) {
	crew := c.Query("crewcount")
	planetid := c.Query("planetid")

	if crew == "" || planetid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing crewcount or planetid parameter"})
		return
	}

	crewCapacity, err := strconv.Atoi(crew)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crewcount parameter"})
		return
	}

	httpClient := http.DefaultClient
	url := fmt.Sprintf("%v/exoplanets/getexoplanet?id=%s", h.config.GetServerConfig().BaseURLExoplanet, planetid)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request", "details": err.Error()})
		return
	}

	res, err := httpClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform request", "details": err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.JSON(res.StatusCode, gin.H{"error": "Failed to fetch exoplanet", "details": http.StatusText(res.StatusCode)})
		return
	}

	var exoplanet store.Exoplanet
	err = json.NewDecoder(res.Body).Decode(&exoplanet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode exoplanet response", "details": err.Error()})
		return
	}

	fuelEstimation := h.fuelEstimatorService.CalculateEstimatedFuel(crewCapacity, exoplanet)
	c.JSON(http.StatusOK, gin.H{"fuelEstimation": fuelEstimation})
}
