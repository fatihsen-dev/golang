package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Person struct {
    Name string `form:"name"`
    Surname string `form:"surname"`
    Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main()  {
	router := gin.Default()

	// allows all origins
	router.Use(cors.Default())

	// GET: localhost:5000?name=Fatih&surname=Şen&birthday=1999-03-18
   router.GET("/", startPage)

	router.Run(":5000")
}

func startPage(c *gin.Context) {
	var person Person
	log.Println(person)
	if c.ShouldBind(&person) == nil {
		log.Printf("Name: %v - Surname: %v - Birthday: %v",person.Name,person.Surname,person.Birthday)
	}

	c.String(200, person.Birthday.Format("2006-01-02"))
}