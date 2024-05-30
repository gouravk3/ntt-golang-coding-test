package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gouravk3/ntt-golang-coding-test/internal/service"
	"github.com/gouravk3/ntt-golang-coding-test/internal/store"
)

type handlers struct {
	exoplanetService *service.ExoplanetService
}

func New() *handlers {
	return &handlers{
		exoplanetService: service.NewExoplanetService(),
	}
}

func (h *handlers) AddExoplanet(c *gin.Context) {
	var exoplanet store.Exoplanet
	err := c.ShouldBindBodyWithJSON(&exoplanet)
	if err != nil {
		log.Println("error while binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.exoplanetService.AddExoplanet(exoplanet)

	c.JSON(http.StatusCreated, gin.H{
		"Added entry": exoplanet,
	})
}

func (h *handlers) ListExoplanets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"List of Exoplanets": h.exoplanetService.ListExoplanets(),
	})
}

func (h *handlers) GetExoplanetByID(c *gin.Context) {
	id := c.Query("id")
	exoplanet, found := h.exoplanetService.GetExoplanetByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": fmt.Sprintf("id: %v, was not found in db", id),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Exoplanet": exoplanet,
	})
}

func (h *handlers) UpdateExoplanet(c *gin.Context) {
	var exoplanet store.Exoplanet
	err := c.ShouldBindBodyWithJSON(&exoplanet)
	if err != nil {
		log.Println("error while binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	updated := h.exoplanetService.UpdateExoplanet(exoplanet)
	if !updated {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "exoplanet not found in db",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Exoplanet": exoplanet,
	})
}

func (h *handlers) DeleteExoplanet(c *gin.Context) {
	id := c.Query("id")
	deleted := h.exoplanetService.DeleteExoplanet(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": fmt.Sprintf("id: %v, was not found in db", id),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "exoplent with was deleted",
	})
}
