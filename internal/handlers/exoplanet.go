package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gouravk3/ntt-golang-coding-test/internal/service"
)

func AddExoplanet(c *gin.Context) {
	var exoplanet service.Exoplanet
	err := c.ShouldBindBodyWithJSON(&exoplanet)
	if err != nil {
		log.Println("error while binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// TODO: add to in-memo db

	c.JSON(http.StatusCreated, exoplanet)
}

func ListExoplanets(c *gin.Context) {
	// TODO: iterate in-memo db and construct a json

	// c.JSON(http.StatusOK, jsonObjext)
	fmt.Println("ListExoplanets")

}

func GetExoplanetByID(c *gin.Context) {
	fmt.Println("GetExoplanetByID")

}

func UpdateExoplanet(c *gin.Context) {

}

func DeleteExoplanet(c *gin.Context) {

}
