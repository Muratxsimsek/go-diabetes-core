package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()
	router.GET("/diabetes", handleGetDiabetesList)
	router.GET("/diabetes/:id", handleGetDiabetes)
	router.POST("/diabetes/", handleCreateDiabetes)
	router.PUT("/diabetes/:id", handleUpdateDiabetes)

	router.Run(":8099")
}

func handleGetDiabetesList(c *gin.Context) {
	var diabetesList, err = GetAllDiabetes()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diabetesList})
}

func handleGetDiabetes(c *gin.Context) {
	id := c.Param("id")
	//if id != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	//	return
	//}
	var savedDiabetes, err = GetDiabetesByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": savedDiabetes})
}

func handleCreateDiabetes(c *gin.Context) {
	var diabetes Diabetes
	if err := c.ShouldBindJSON(&diabetes); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := CreateDiabetes(&diabetes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func handleUpdateDiabetes(c *gin.Context) {
	id := c.Param("id")
	var diabetes Diabetes
	if err := c.ShouldBindJSON(&diabetes); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	upsertedID, err := UpdateDiabetes(id, &diabetes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": upsertedID})
}
