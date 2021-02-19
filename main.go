package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func handleGetDiabetes(c *gin.Context) {
//	var diabetesList []Diabetes
//	var diabetes Diabetes
//	diabetes.HungerStatus = "FASTING"
//	diabetes.SugarValue = 123
//
//	diabetesList = append(diabetesList, diabetes)
//	c.JSON(http.StatusOK, gin.H{"all diabetes": diabetesList})
//}

func handleGetDiabetes(c *gin.Context) {
	var diabetesList, err = GetAllDiabetes()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"diabetes": diabetesList})
}

func main() {

	router := gin.Default()
	router.GET("/diabetes", handleGetDiabetes)
	//router.POST("/users", controllers.UsersController.Save)

	router.Run(":8099")
}
