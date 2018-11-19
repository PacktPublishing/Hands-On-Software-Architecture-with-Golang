package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

// User desscribes an employee entity.
// We will be using this to demonstrate Cassandra create, read, update, and delete flows
type User struct {
	Id        gocql.UUID
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "roster"
	session, _ := cluster.CreateSession()
	defer session.Close()

	// generate a unique id for the employee
	id := gocql.TimeUUID()
	// create the employee in memory
	newEmployee := User{
		Id:        id,
		FirstName: "James",
		LastName:  "Bond",
		Age:       45,
	}

	// insert the employee
	if err := session.Query("INSERT INTO employees (id, firstname, lastname, age ) VALUES (?, ?, ?, ?)",
		newEmployee.Id,
		newEmployee.FirstName,
		newEmployee.LastName,
		newEmployee.Age).Exec(); err != nil {
		fmt.Println("insert error")
		log.Fatal(err)
	}

	// Use select to get the employee we just entered
	var userFromDB User

	if err := session.Query("SELECT id, firstname,  lastname, age FROM employees WHERE id=?", id).
		Scan(&userFromDB.Id, &userFromDB.FirstName, &userFromDB.LastName, &userFromDB.Age); err != nil {
		fmt.Println("select error")
		log.Fatal(err)
	}
	fmt.Println(userFromDB)

	// Update James's Bond's age
	if err := session.Query("UPDATE employees SET age = 46 WHERE id = ?", id).Exec(); err != nil {
		fmt.Println("udpate error")
		log.Fatal(err)
	}

	var newAge int
	// Select and show the change
	iter := session.Query("SELECT age FROM employees WHERE  id = ?", id).Iter()
	for iter.Scan(&newAge) {
		fmt.Println(newAge)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	// Delete the employe
	if err := session.Query("DELETE FROM employees WHERE id = ?", id).Exec(); err != nil {
		fmt.Println("delete error")
		log.Fatal(err)
	}
}
