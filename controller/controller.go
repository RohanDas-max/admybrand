package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rohandas-max/admybrand/database"
	"github.com/rohandas-max/admybrand/model"
)

func GetData(c *gin.Context) {
	var model *[]model.Data
	c.Writer.Header().Set("Content-Type", "application/json")
	if err := database.DB.Find(&model).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": model,
		})
	}

}

func GetAData(c *gin.Context) {
	var model *model.Data
	userId := c.Param("id")
	c.Writer.Header().Set("Content-Type", "application/json")

	if err := database.DB.Where("id =? ", userId).First(&model).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": model,
		})
	}

}

func CreateData(c *gin.Context) {
	var model *model.Data
	c.Header("Content-Type", "application/json")
	if err := c.Bind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
			"err":   err,
		})
	} else {
		err := database.DB.Create(&model).Error
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"err": "cannot Insert", "msg": err,
			})
		} else {
			c.JSON(http.StatusAccepted, gin.H{
				"success": "Data Created", "data": model,
			})
		}
	}
}

func UpdateData(c *gin.Context) {
	var model *model.Data
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)
	c.Writer.Header().Set("Content-Type", "application/json")

	if err := c.Bind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
			"err":   err,
		})
	}

	err := database.DB.Where("id = ?", id).Save(&model).Error

	// 	"UPDATE table_name
	// SET column1 = value1, column2 = value2...., columnN = valueN
	// WHERE [condition];"

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"err":   "an error occured check",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":     "Data updated",
			"updated": model,
		})
	}

}

func DeleteData(c *gin.Context) {

	var model *model.Data
	userId := c.Param("id")
	c.Writer.Header().Set("Content-Type", "application/json")

	if err := database.DB.Where("id = ?", userId).Delete(&model).Error; err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"err":   "please enter a valid id ",
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Data deletd",
		})
	}

}
