package controllers

import (
	"encoding/json"
	"fmt"
	"net/http" //  help in working with http requests in golang
	"time"

	"github.com/inspo13/aInstagram-Api/models"

	"github.com/julienschmidt/httprouter"

	"gopkg.in/mgo.v2" //mongodb package
	"gopkg.in/mgo.v2/bson"
)

var PostedTimestamp string

type PostController struct {
	session *mgo.Session
}

func NewPostController(s *mgo.Session) *PostController {
	return &PostController{s}
}

//create user method for creating post
func (pc PostController) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	PostedTimestamp := time.Now()
	pt := models.Post{}
	//decode decodes json to bson
	json.NewDecoder(r.Body).Decode(&pt)
	pt.Id = bson.NewObjectId()
	pc.session.DB("aInstagram-Api").C("posts").Insert(pt)
	uj, err := json.Marshal(pt)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
	fmt.Println(PostedTimestamp.String())

}

//get  post method for reading the post
func (pc PostController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	if bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.IsObjectIdHex(id)
	pt := models.Post{}
	if err := pc.session.DB("aInstagram-Api").C("posts").FindId(oid).One(&pt); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

//get all the posts created by the user
func (pc PostController) GetAllPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//here the  id id from user details
	id := p.ByName("id")
	if bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.IsObjectIdHex(id)
	pt := models.Post{}
	if err := pc.session.DB("aInstagram-Api").C("posts/users").FindId(oid).One(&pt); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}
