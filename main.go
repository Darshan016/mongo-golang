package main

import (
	"net/http"

	"log"

	"github.com/Darshan016/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Println("Server running on port: 8000")
	http.ListenAndServe("localhost:8000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("error while connecting to db: %v", err)
	}
	return s
}
