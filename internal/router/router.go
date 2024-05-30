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

	h := handlers.New()

	apiGroup := r.Group("/exoplanets/api/v1")
	healthRouter(apiGroup)
	exoplanetGroup := apiGroup.Group("/exoplanets")
	{
		exoplanetGroup.POST("", h.AddExoplanet)
		exoplanetGroup.GET("", h.ListExoplanets)
		exoplanetGroup.GET("/getexoplanet", h.GetExoplanetByID)
		exoplanetGroup.PUT("/updateexoplanet", h.UpdateExoplanet)
		exoplanetGroup.DELETE("/deleteexoplanet", h.DeleteExoplanet)
	}

	return r, nil
}

func FuelEstimationRouter(appConfig config.AppConfig) (*gin.Engine, error) {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	handler := handlers.NewfuelEstimatorHandler(appConfig)

	apiGroup := r.Group("/fuelestimation/api/v1")
	healthRouter(apiGroup)
	apiGroup.GET("/fuelestimation", handler.FuelEstimation)

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
