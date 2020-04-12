package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Thing is a struct that holds fields for a person, their age, and favorite food
type Thing struct {
	Person       string
	Age          int
	FavoriteFood string
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

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/person/:name", func(c *gin.Context) {
		name := c.Param("name")
		param := getPerson(name, db)
		c.JSON(200, param)
	})

	r.PUT("/person/create", func(c *gin.Context) {
		var newPerson Thing
		c.BindJSON(&newPerson)
		putPerson(newPerson, db)
		c.JSON(201, newPerson)
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

	r.PATCH("/person/:name", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{"status": "not implemented"})
	})

	r.GET("/persons/list", func(c *gin.Context) {
		persons := getPersons(db)
		c.JSON(http.StatusOK, persons)
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

func getPersons(db *sql.DB) (persons []Thing) {
	rows, err := db.Query(`SELECT Age, Person, FavoriteFood FROM helloworld.person`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var person Thing
		err := rows.Scan(&person.Age, &person.Person, &person.FavoriteFood)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, person)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func deletePerson(name string, db *sql.DB) bool {
	_, err := db.Exec(`DELETE FROM helloworld.person WHERE person.Person = $1`, name)
	if err != nil {
		return false
	}
	return true
}
