package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type myForm struct {
    Name string `form:"name"`
    Surname string `form:"surname"`
    Age int `form:"age"`
	 Skils []string `form:"skils[]"`
}

func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)
    c.JSON(200, gin.H{
		"name": fakeForm.Name,
		"surname": fakeForm.Surname,
		"age": fakeForm.Age,
		"skils": fakeForm.Skils,
	})
}

func main()  {
	router := gin.Default()

	// allows all origins
	router.Use(cors.Default())
   router.POST("/", formHandler)

	router.Run(":5000")
}