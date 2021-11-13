package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohandas-max/admybrand/controller"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.GET("/api/getdata", controller.GetData)
	r.GET("/api/:id/getadata", controller.GetAData)
	r.POST("/api/createdata", controller.CreateData)
	r.PUT("/api/:id/updatedata", controller.UpdateData)
	r.DELETE("/api/:id/deletedata", controller.DeleteData)

	return r
}
