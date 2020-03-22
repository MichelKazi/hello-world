package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Thing struct {
	Person       string
	Age          int
	FavoriteFood string
}

type Database struct {
	People []Thing
}

func main() {
	db := Database{}
	db.People = append(db.People, Thing{Person: "Hillary", Age: 27, FavoriteFood: "Toads"})
	db.People = append(db.People, Thing{Person: "Jeffry", Age: 37, FavoriteFood: "Snickers"})
	db.People = append(db.People, Thing{Person: "Tyler", Age: 25, FavoriteFood: "Flies"})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hillary",
		})
	})

	r.GET("/hillary/:person", func(c *gin.Context) {
		person := c.Param("person")
		found := false
		for _, v := range db.People {
			if v.Person == person {
				found = true
				c.JSON(200, v)
			}
		}
		if !found {
			c.JSON(404, gin.H{"Error": "Nothing found"})
		}
	})

	r.GET("/people", func(c *gin.Context) {
		c.JSON(200, db.People)
	})

	r.PUT("/person/create", func(c *gin.Context) {
		var person Thing
		c.BindJSON(&person)
		fmt.Println(person.Person)
		c.JSON(201, person)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
