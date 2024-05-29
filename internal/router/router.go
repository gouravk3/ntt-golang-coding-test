package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gouravk3/ntt-golang-coding-test/config"
)

type HealthController struct{}

func InitRouter(appConfig config.AppConfig) (*gin.Engine, error) {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	health := new(HealthController)

	r.GET("/health", health.Status)
	r.POST("/exoplanets", AddExoplanet)
	r.GET("/exoplanets", ListExoplanets)
	r.GET("/exoplanets/{id}", GetExoplanetByID)
	r.PUT("/exoplanets/{id}", UpdateExoplanet)
	r.DELETE("/exoplanets/{id}", DeleteExoplanet)
	r.GET("/exoplanets/{id}/fuel/", FuelEstimation)

	return r, nil
}

func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Still Working!")
}

func AddExoplanet(c *gin.Context) {

}

func ListExoplanets(c *gin.Context) {

}

func GetExoplanetByID(c *gin.Context) {

}

func UpdateExoplanet(c *gin.Context) {

}

func DeleteExoplanet(c *gin.Context) {

}

func FuelEstimation(c *gin.Context) {

}
