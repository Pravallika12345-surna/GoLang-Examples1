package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rizalgowandy/go-swag-sample/database"
	_ "github.com/rizalgowandy/go-swag-sample/docs"
	"github.com/rizalgowandy/go-swag-sample/model"
	"github.com/rizalgowandy/go-swag-sample/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var err error

func main() {
	database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer database.DB.Close()
	database.DB.AutoMigrate(&model.Employee{})
	//r := gin.New()
	r := routes.Setuprouter()

	r.GET("/emp", EmpGet)
	r.POST("/emp", EmpCreate)
	r.GET("/emp/{id}", GetEmpByID)
	r.PUT("/emp/{id}", EmpUpdate)
	r.DELETE("/emp/{id}", EmpDelete)

	url := ginSwagger.URL("http://localhost:8060/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if err := r.Run(":8060"); err != nil {
		log.Fatal(err)
		r.Run()
	}
}

func EmpGet(c *gin.Context) {
	var emp []model.Employee
	err := model.GetAllEmp(&emp)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

func EmpCreate(c *gin.Context) {
	var emp model.Employee
	c.BindJSON(&emp)
	err := model.CreateEmp(&emp)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

func GetEmpByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var emp model.Employee
	err := model.GetEmpByID(&emp, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

func EmpUpdate(c *gin.Context) {
	var emp model.Employee
	id := c.Params.ByName("id")
	err := model.GetEmpByID(&emp, id)
	if err != nil {
		c.JSON(http.StatusNotFound, emp)
	}
	c.BindJSON(&emp)
	err = model.EmpUpdate(&emp, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

func EmpDelete(c *gin.Context) {
	var emp model.Employee
	id := c.Params.ByName("id")
	err := model.EmpDelete(&emp, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
