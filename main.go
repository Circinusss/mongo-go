package main

import (
	"net/http"

	"github.com/Circinusss/mongo-go/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.GET("/posts/:id", pc.GetPost)
	r.POST("/posts", pc.CreatePost)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
