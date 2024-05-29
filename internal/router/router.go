package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gouravk3/ntt-golang-coding-test/config"
	"github.com/gouravk3/ntt-golang-coding-test/internal/handlers"
)

type HealthController struct{}

func ExoplanetRouter(appConfig config.AppConfig) (*gin.Engine, error) {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiGroup := r.Group("/exoplanets/api/v1")
	healthRouter(apiGroup)
	exoplanetGroup := apiGroup.Group("/exoplanets")
	{
		exoplanetGroup.POST("", handlers.AddExoplanet)
		exoplanetGroup.GET("", handlers.ListExoplanets)
		exoplanetGroup.GET("/:id", handlers.GetExoplanetByID)
		exoplanetGroup.PUT("/:id", handlers.UpdateExoplanet)
		exoplanetGroup.DELETE("/:id", handlers.DeleteExoplanet)
	}

	return r, nil
}

func FuelEstimationRouter(appConfig config.AppConfig) (*gin.Engine, error) {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiGroup := r.Group("/fuelestimation/api/v1")
	healthRouter(apiGroup)
	apiGroup.POST("/fuelestimation", handlers.FuelEstimation)

	return r, nil
}

func healthRouter(apiGroup *gin.RouterGroup) {
	health := new(HealthController)

	healthGroup := apiGroup.Group("/health")
	{
		healthGroup.GET("", health.Status)
	}
}

func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Still Working!")
}
