package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Thing struct {
	Person       string
	Age          int
	FavoriteFood string
}

type Database struct {
	People []Thing
}

const (
	host     = "fullstack-postgres"
	port     = 5432
	user     = "admin"
	password = "admin123"
	dbname   = "dev"
)

func main() {
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	log.Printf("Postgres started at %d PORT", port)
	defer db.Close()

	data := Database{}
	data.People = append(data.People, Thing{Person: "Hillary", Age: 27, FavoriteFood: "Toads"})
	data.People = append(data.People, Thing{Person: "Jeffry", Age: 37, FavoriteFood: "Snickers"})
	data.People = append(data.People, Thing{Person: "Tyler", Age: 25, FavoriteFood: "Flies"})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hillary",
		})
	})

	r.GET("/hillary/:person", func(c *gin.Context) {
		person := c.Param("person")
		found := false
		for _, v := range data.People {
			if v.Person == person {
				found = true
				c.JSON(200, v)
			}
		}
		if !found {
			c.JSON(404, gin.H{"Error": "Nothing found"})
		}
	})

	r.GET("/people/:name", func(c *gin.Context) {
		name := c.Param("name")
		person := getPerson(name, db)
		c.JSON(200, person)
	})

	r.PUT("/person/create", func(c *gin.Context) {
		var person Thing
		c.BindJSON(&person)
		putPerson(person, db)
		c.JSON(201, person)
	})

	r.DELETE("/person/:name", func(c *gin.Context) {
		name := c.Param("name")
		result := deletePerson(name, db)
		if result {
			c.JSON(200, gin.H{"Person": name, "Status": "deleted"})
		} else {
			c.JSON(400, gin.H{"Person": name, "Status": "not-found"})
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func putPerson(person Thing, db *sql.DB) {
	_, err := db.Exec(`INSERT into helloworld.person (Person, Age, FavoriteFood) VALUES ($1, $2, $3)`, person.Person, person.Age, person.FavoriteFood)
	if err != nil {
		panic(err)
	}
}

func getPerson(name string, db *sql.DB) (person Thing) {
	row := db.QueryRow(`SELECT Age, Person, FavoriteFood FROM helloworld.person WHERE Person = $1`, name)
	err := row.Scan(&person.Age, &person.Person, &person.FavoriteFood)

	switch err {
	case sql.ErrNoRows:
		return
	case nil:
		return
	default:
		panic(err)
	}
}

func deletePerson(name string, db *sql.DB) bool {
	_, err := db.Exec(`DELETE FROM helloworld.person WHERE person.Person = $1`, name)
	if err != nil {
		return false
	}
	return true
}
