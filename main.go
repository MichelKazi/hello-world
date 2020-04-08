package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Thing is a struct that holds fields for a person, their age, and favorite food
type Thing struct {
	Person       string
	Age          int
	FavoriteFood string
}

// Database is a struct containing a slice of people
type Database struct {
	People []Thing
}

// Person is a struct composed of a peron's name and age
type Person struct {
	Name string
	Age  int
}

// GetName returns the name of a Person object
func (p Person) GetName() string {
	return p.Name
}

// SetName takes a string and sets that as the name of a Person object
func (p *Person) SetName(realName string) {
	p.Name = fmt.Sprintf("%s - Age", realName)
}

// GetAge returns the age of a Person object
func (p Person) GetAge() int {
	return p.Age
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

	person := Person{}
	person.Name = "hillary"
	person.Age = 27
	person.SetName("Phil")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": person.GetName(),
			"age":  person.GetAge(),
		})
	})

	r.GET("/hillary/:person", func(c *gin.Context) {
		param := c.Param("person")
		found := false
		for _, v := range data.People {
			if v.Person == param {
				found = true
				c.JSON(200, v)
			}
		}
		if !found {
			c.JSON(404, gin.H{"Error": "Nothing found"})
		}
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

	r.PATCH("/person/update/:name", func(c *gin.Context) {
		var newNamePerson Thing
		c.BindJSON(&newNamePerson)
		name := c.Param("name")
		//newName := c.Param("newName")
		newName := newNamePerson.Person

		patchPersonName(name, newName, db)

		param := getPerson(newName, db)

		c.JSON(204, param)
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
	fmt.Println(row)

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

func patchPersonName(name string, newName string, db *sql.DB) {
	_, err := db.Exec(`UPDATE helloworld.person SET Person = $1 WHERE Person = $2`, newName, name)
	if err != nil {
		panic(err)
	}
}
