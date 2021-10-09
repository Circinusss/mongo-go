package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Circinusss/mongo-go/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostController struct {
	session *mgo.Session
}

func NewPostController(s *mgo.Session) *PostController {
	return &PostController{s}
}

func (pc PostController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.Post{}

	if err := pc.session.DB("mongo-golang").C("posts").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	pj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", pj)
}

func (pc PostController) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.Post{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	pc.session.DB("mongo-golang").C("posts").Insert(u)

	pj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", pj)
}
